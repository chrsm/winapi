winapi
======

[![GoDoc](https://godoc.org/github.com/chrsm/winapi?status.svg)](https://godoc.org/github.com/chrsm/winapi)

This is a work-in-progress. I'm only adding things that I need right now.
Some stuff is copied from other projects I have and as such may not actually function.

I wanted to use some Windows APIs, but in a fashion that didn't make me want to tear my hair out.
So I wrote some wrappers for Windows APIs to make them feel "like Go" so I don't have to
tear my hair out later.

Since these APIs differ from their "native" counterparts in that they're a bit more idiomatic,
I recommend reading the godoc for each package.


Supported APIs
==============

- user: APIs from user32.dll / winuser.h

- wslapi:
  - wsl.ConfigureDistribution = WslConfigureDistribution
  - wsl.GetDistributionConfiguration = WslGetDistributionConfiguration
  - wsl.DistributionRegistered = WslDistributionRegistered
  - wsl.Launch = WslLaunch
  - wsl.LaunchInteractive = WslLaunchInteractive
  - wsl.RegisterDistribution = WslRegisterDistribution
  - wsl.UnregisterDistribution = WslUnregisterDistribution

