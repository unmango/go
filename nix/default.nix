{
  pname,
  buildGoApplication,
  ginkgo,
  git,
  go,
  lib,
  version,
}:
buildGoApplication {
  inherit go pname version;

  src = lib.cleanSource ../.;
  modules = ./gomod2nix.toml;
  nativeBuildInputs = [ ginkgo ];
  nativeCheckInputs = [ git ];

  # vcs/git's tests expect the source root to be a git repository.
  checkPhase = ''
    runHook preCheck
    git init --quiet
    make test
    runHook postCheck
  '';
}
