class Dt < Formula
  desc "CLI tool for common dev tasks."
  homepage "https://github.com/chclaus/dt"
  url "https://github.com/chclaus/dt/releases/download/v1.1.0/dt_1.1.0_darwin_amd64.tar.gz"
  version "1.1.0"
  sha256 "aa03d4126a9fe1ef96e55c25c092f0da3d50fafa09a40dd92f5e45787d46f93e"

  def install
    bin.install "dt"
    zsh_completion.install "dt-completion.zsh" => "_dt.zsh"
    bash_completion.install "dt-completion.bash"
  end
end
