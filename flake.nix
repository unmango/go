{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
		flake-utils.url = "github:numtide/flake-utils";
		treefmt-nix = {
			url = "github:numtide/treefmt-nix";
			inputs.nixpkgs.follows = "nixpkgs";
		};
  };

  outputs = { self, nixpkgs, flake-utils, treefmt-nix }: flake-utils.lib.eachDefaultSystem (system:
		let
			pkgs = nixpkgs.legacyPackages.${system};
		in
		{
			formatter = treefmt-nix.lib.mkWrapper pkgs {
				projectRootFile = "flake.nix";
				programs.gofmt.enable = true;
			};

			devShells.default = pkgs.mkShell {
				buildInputs = with pkgs; [
					git
					gnumake
					go
					ginkgo
				];
			};
		}
	);
}
