class Dt < Formula
  desc "CLI tool for common dev tasks."
  homepage "https://github.com/chclaus/dt"
  url "https://github.com/chclaus/dt/releases/download/v0.2.0/dt_0.2.0_darwin_amd64.tar.gz"
  version "0.2.0"
  sha256 "5d200976656c5f3f04216f65c987ef802fb730275843cb4d7479c98647ab59d4"

  def install
    bin.install "dt"
  end
end
