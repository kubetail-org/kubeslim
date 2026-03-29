# Contributing to kubeslim

Thank you for your interest in contributing to kubeslim! We're building an alternative to client-go optimized for binary size, memory consumption, and speed, and we'd love your help.

This document will guide you through the contribution process.

## Table of Contents

- [Where to Find Code](#where-to-find-code)
- [How to Run Tests and Other Checks](#how-to-run-tests-and-other-checks)
- [Commit Guidelines](#commit-guidelines)
- [Pull Request Guidelines](#pull-request-guidelines)
- [Branch Naming Guidelines](#branch-naming-guidelines)
- [Editor Configuration](#editor-configuration)
- [Automation](#bots-and-automation)
- [AI Policy](#ai-policy)
- [Community](#community)

## Where to Find Code

kubeslim is a single Go module at the repository root:

| File / Directory | Purpose               |
|------------------|-----------------------|
| `client.go`      | Client implementation |
| `resource.go`    | Resource helpers      |
| `examples/`      | Usage examples        |

## How to Run Tests and Other Checks

Make sure your changes pass all tests and checks before submitting a pull request.

```bash
# Format check
test -z $(gofmt -l .)

# Vet (static analysis)
go vet ./...

# Run tests
go test -race ./...
```

## Commit Guidelines

All commits should be squashed into a single, signed commit before merging.

### Format

We follow the [Conventional Commits](https://www.conventionalcommits.org) format:

```
<type>[optional scope]: commit title goes here (all lowercase)

[optional body]

Signed-off-by: Your Name <you@example.com>
```

**Types:**

- `build` - Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
- `ci` - Changes to our CI configuration files and scripts
- `docs` - Documentation only changes
- `feat` - A new feature
- `fix` - A bug fix
- `perf` - A code change that improves performance
- `refactor` - A code change that neither fixes a bug nor adds a feature
- `style` - Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- `test` - Adding missing tests or correcting existing tests

## Pull Request Guidelines

### Before Submitting

1. **Check for duplicates**: Review existing [issues](https://github.com/kubetail-org/kubeslim/issues) and [pull requests](https://github.com/kubetail-org/kubeslim/pulls)
2. **Run tests**: Execute all relevant tests for your changes and ensure they pass
3. **Format code**: Run `gofmt` on any modified Go files
4. **Update branch**: Rebase your branch to the latest `main`
5. **Squash commits**: Combine all commits into a single, signed commit following our [commit format](https://www.conventionalcommits.org)

### PR Title Format

Add an emoji to indicate the PR type:

- 🎣 Bug fix
- 🐋 New feature
- 📜 Documentation
- ✨ General improvement

### PR Description

Your PR should include:

- Link to related issue: `Fixes #123`
- **Summary**: Explain the goal of your PR
- **Key Changes**: List the specific key changes made

### PR Checklist

- [ ] Add the correct emoji to the PR title
- [ ] Link the issue number, if any, to `Fixes #`
- [ ] Add summary and explain key changes
- [ ] Rebase branch to HEAD
- [ ] Squash changes into one signed, single commit

## Branch Naming Guidelines

Use descriptive branch names with this pattern:

```
<type>/<short-description>
```

## Editor Configuration

### Visual Studio Code

Recommended extensions:
- **Go**: `golang.go`

## Bots and Automation

### GitHub Actions

Our CI/CD pipeline automatically runs on every pull request:

- **Tests**: All unit and integration tests
- **Linting**: Code formatting and style checks
- **Build**: Ensures the module builds successfully

You can see the status of these checks in your PR. If any checks fail, review the logs and fix the issues before requesting a review.

### CLA Assistant

If this is your first contribution, our [CLA (Contributor License Agreement)](https://cla-assistant.io/) assistant will prompt you to sign the CLA when you create your pull request. This is a one-time requirement.

## AI Policy

As a contributor you're encouraged to use AI tools in your workflow just as you would use classic tools such as search engines, language servers, linters, debuggers, documentation, or books. These tools are an invaluable resource and can help you write better code and explore ideas more efficiently.

That said, AI tools are different than classic tools because they can blur the line between helping you to do the work and doing the work for you. And when that line becomes blurry, it can limit opportunities to build the deep understanding that comes from writing and reasoning through code yourself.

As an open source project, kubeslim is not only committed to building a lean, high-performance Kubernetes client library but also to helping our contributors grow as engineers. We invest a lot of time and effort into code quality, thoughtful reviews, and well-defined engineering specs. We do so happily because we enjoy it but also because it's our responsibility to the community.

In return, we ask that contributions be authored by you. While AI tools can support your workflow, submitted code should reflect your own understanding and intent. To keep our focus on meaningful collaboration within the community, we do not accept contributions generated by bots or submissions authored primarily by llms.

## Community

We'd love to hear from you! Here's how to connect with the kubeslim community.

### Communication Channels

- **[Discord](https://discord.gg/CmsmWAVkvX)**: Join for real-time discussions, questions, and community chat
- **[Slack](https://kubernetes.slack.com/archives/C08SHG1GR37)**: Connect with us on the Kubernetes workspace

### Code of Conduct

Please read and follow our [Code of Conduct](https://github.com/kubetail-org/.github/blob/main/CODE_OF_CONDUCT.md). We are committed to providing a welcoming and inclusive environment for all contributors.

---

Thank you for contributing to kubeslim!
