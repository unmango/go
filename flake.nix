{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs = {
        nixpkgs.follows = "nixpkgs";
        flake-utils.follows = "flake-utils";
      };
    };
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      treefmt-nix,
      gomod2nix,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};

        go-test = pkgs.stdenvNoCC.mkDerivation {
          name = "go-test";
          src = ./.;
          doCheck = true;
          nativeBuildInputs = with pkgs; [
            git
            go
            ginkgo
            writableTmpDirAsHomeHook
          ];
          # checkPhase = ''
          #   go test ./... -ginkgo.label-filter="Dependency: isEmpty"
          # '';
          checkPhase = ''
            ginkgo run -r --label-filter="Dependency: isEmpty"
          '';
          installPhase = ''
            mkdir "$out"
          '';
        };
      in
      {
        formatter = treefmt-nix.lib.mkWrapper pkgs {
          projectRootFile = "flake.nix";
          programs.gofmt.enable = true;
          programs.nixfmt.enable = true;
        };

        checks = {
          inherit go-test;
        };

        packages.default = pkgs.callPackage ./. {
          inherit (gomod2nix.legacyPackages.${system}) buildGoApplication;
        };

        devShells.default = pkgs.callPackage ./shell.nix {
          inherit (gomod2nix.legacyPackages.${system}) mkGoEnv gomod2nix;
        };
      }
    );
}
