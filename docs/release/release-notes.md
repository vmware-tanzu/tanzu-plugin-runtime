# Release notes

This document provides guidance on providing release notes for changes made to
Tanzu Plugin Runtime. Release notes act as a direct line of communication
with the main consumers of the runtime -- the plugin developers. In particular
they should also encompass information that could result in user-facing changes
in the plugins, since those changes might have to in turn be communicated to
the plugin users.

## Table of Contents

* [Does my pull request need a release note?](#does-my-pull-request-need-a-release-note)
* [Contents of a release note](#contents-of-a-release-note)
* [Applying a Release Note](#applying-a-release-note)
* [Reviewing Release Notes](#reviewing-release-notes)

## Does my pull request need a release note?

Any pull request that could impact a runtime consumer's decision in picking up
a new release is required to add release notes. This could mean:

* Critical bug-fixes
* Notable feature additions
* Output format changes
* API changes, including deprecation and removals
* Configuration schema change

No release notes are required for changes to:

* Tests
* Build infrastructure and general repository maintenance

## Contents of a release note

A release note needs a clear, concise description of the change in simple plain language.
This includes:

* An indicator if the pull request Added, Changed, Fixed, Removed, Deprecated functionality or changed enhancement.
* The name of the affected API, configuration field or affected component.
* A link to relevant user documentation about the enhancement/feature.

Your release note should be written in clear and straightforward sentences.
Not all users are familiar with the technical details of your pull request,
so consider what they need to know when you write your release note.

Write the release note like you are in the future like:

* "Added" instead of "add"
* "Fixed" instead of "fix"
* "Bumped" instead of "bump"

Some examples of release notes:

* API Foo has been deprecated, will be removed in version "1.5.0".
  Users of said API should use "Bar" instead.
* Fixed a bug that prevents the configuration from being read from shared folder.

## Applying a Release Note

Any pull request with user visible changes should include a release-note block in the pull request description.

For pull requests with a release note:

```text
    ```release-note
    Your release note here
    ```
```

For pull requests with no release note:

```text
    ```release-note
    NONE
    ```
```

## Reviewing Release Notes

Release note should be reviewed as a dedicated step in the overall pull request
review process.

A release note needs to be changed or rephrased if one of the following cases
apply:

* The release note does not communicate the full purpose of the change.
* The release note is grammatically incorrect.
* The release does not comply with the guidelines on the contents of the release note.
