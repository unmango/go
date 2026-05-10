{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    systems.url = "github:nix-systems/default";
    flake-parts.url = "github:hercules-ci/flake-parts";

    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import inputs.systems;
      imports = [ inputs.treefmt-nix.flakeModule ];

      perSystem =
        {
          inputs',
          pkgs,
          lib,
          system,
          ...
        }:
        let
          inherit (inputs'.gomod2nix.legacyPackages) buildGoApplication;
        in
        {
          _module.args.pkgs = import inputs.nixpkgs {
            inherit system;
            overlays = [ inputs.gomod2nix.overlays.default ];
          };

          packages.default = pkgs.callPackage ./nix {
            inherit buildGoApplication;
            pname = "go";
            version = "0.15.1";
          };

          devShells.default = pkgs.mkShell {
            buildInputs = with pkgs; [
              git
              gnumake
              go
              gomod2nix
              gopls
              ginkgo
              nixfmt
            ];

            GINKGO = "${pkgs.ginkgo}/bin/ginkgo";
            GO = "${pkgs.go}/bin/go";
            GOMOD2NIX = "${pkgs.gomod2nix}/bin/gomod2nix";
          };

          treefmt = {
            programs.actionlint.enable = true;
            programs.nixfmt.enable = true;
            programs.gofmt.enable = true;
          };
        };
    };
}
