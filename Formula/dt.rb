class Dt < Formula
  desc "CLI tool for common dev tasks."
  homepage "https://github.com/chclaus/dt"
  url "https://github.com/chclaus/dt/releases/download/v0.1.2-beta/dt_0.1.2-beta_darwin_amd64.tar.gz"
  version "0.1.2-beta"
  sha256 "738c935e30addb931c956e8f0e75281f07db5ae67f113bbe6932b790b7692cb1"

  def install
    bin.install "dt"
  end
end
