with import <nixpkgs> { };
pkgs.mkShell {
  buildInputs =
    [ glib pkgconfig git rustup gnumake protobuf clojure go_1_16 nmap ];
}
