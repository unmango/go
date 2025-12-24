{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    systems.url = "github:nix-systems/default";
    flake-utils.url = "github:numtide/flake-utils";
    flake-parts.url = "github:hercules-ci/flake-parts";

    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.follows = "flake-utils";
    };
  };

  outputs =
    inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import inputs.systems;

      imports = [ inputs.treefmt-nix.flakeModule ];

      perSystem =
        { inputs', pkgs, system, ... }:
        let
          inherit (inputs'.gomod2nix.legacyPackages) buildGoApplication mkGoEnv;

          mangoGo = buildGoApplication {
            pname = "go";
            version = "0.10.2";
            src = ./.;
            modules = ./gomod2nix.toml;

            nativeBuildInputs = [ pkgs.ginkgo ];

            checkPhase = ''
              ginkgo -r --label-filter="Dependency: isEmpty"
            '';
          };
        in
        {
          _module.args.pkgs = import inputs.nixpkgs {
            inherit system;
            overlays = [ inputs.gomod2nix.overlays.default ];
          };

          packages.mangoGo = mangoGo;
          packages.default = mangoGo;

          devShells.default = pkgs.mkShell {
            buildInputs = with pkgs; [
              git
              gnumake
              go
              gomod2nix
              ginkgo
              nixfmt
            ];
          };

          treefmt = {
            programs.nixfmt.enable = true;
            programs.gofmt.enable = true;
          };
        };
    };
}
