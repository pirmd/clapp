# Changelog
All notable changes to this project will be documented in this file.

Format is based on [Keep a Changelog] (https://keepachangelog.com/en/1.0.0/).
Versionning adheres to [Semantic Versioning] (https://semver.org/spec/v2.0.0.html)

## [0.3.2] - 2020-02-21
### Removed
- fully get rid of depreciated github.com/pirmd/cli

## [0.3.1] - 2019-11-11
### Removed
- separate app from rest of github.com/pirmd/cli repository

## [0.3.0] - 2019-11-10
### Added
- now support basic config files management
### Modified
- refactor code to allow direct definition of Commands and Apps (See Examples)

## [0.2.0] - 2019-08-11
### Added
- add function to generate a help file in markdown format for a command 
- add default support to print a version information taken from 'git
  describe and rev-parse' and set-up using 'ldflags -X' directive. Provides a
  simple shell script to facilitate the build/install directive for that purpose.
  This behaviour can be overwriten as wished.
### Modified
- allow command exxecution if no args have been specified by user

## [0.1.0] - 2019-05-11
### Added
- commandline app definition with flags and args parsing and help and/or
  manpage generation
