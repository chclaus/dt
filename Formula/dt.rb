class Dt < Formula
  desc "CLI tool for common dev tasks."
  homepage "https://github.com/chclaus/dt"
  url "https://github.com/chclaus/dt/releases/download/v1.0.5/dt_1.0.5_darwin_amd64.tar.gz"
  version "1.0.5"
  sha256 "37b07c51148d98f7ca1bce2de40a257e301dd13e8d4c7fe9c485c894b51003bd"

  def install
    bin.install "dt"
    zsh_completion.install "dt-completion.zsh" => "_dt.zsh"
    bash_completion.install "dt-completion.bash"
  end
end
