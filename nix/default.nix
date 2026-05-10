{
  pname,
  buildGoApplication,
  ginkgo,
  lib,
  version,
}:
buildGoApplication {
  inherit pname version;

  src = lib.cleanSource ../.;
  modules = ./gomod2nix.toml;
  nativeBuildInputs = [ ginkgo ];

  checkPhase = ''
    ginkgo -r --label-filter="Dependency: isEmpty"
  '';
}
