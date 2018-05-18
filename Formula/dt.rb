class Dt < Formula
  desc "CLI tool for common dev tasks."
  homepage "https://github.com/chclaus/dt"
  url "https://github.com/chclaus/dt/releases/download/v1.0.2/dt_1.0.2_darwin_amd64.tar.gz"
  version "1.0.2"
  sha256 "04e3344beaa5b8f5f678b1e8292674a926a73189e497c81eeb69ed7c4f8c45d8"

  def install
    bin.install "dt"
  end
end
