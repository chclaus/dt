class Dt < Formula
  desc "CLI tool for common dev tasks."
  homepage "https://github.com/chclaus/dt"
  url "https://github.com/chclaus/dt/releases/download/v1.0.0/dt_1.0.0_darwin_amd64.tar.gz"
  version "1.0.0"
  sha256 "380e4937c428f35fc7a55a2ab6d7db8eefcbfba2a7993b24694f3f8348d9ab39"

  def install
    bin.install "dt"
  end
end
