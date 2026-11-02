<div align="center">

# gh-attach

[![version](https://badgen.net/github/release/Sygix/gh-attach?label=version)](https://github.com/Sygix/gh-attach/releases)
[![license](https://badgen.net/github/license/Sygix/gh-attach?color=green)](./LICENSE)

A GitHub CLI extension that uploads and downloads GitHub attachments.

Fork of [sudosubin/gh-attach](https://github.com/sudosubin/gh-attach) by Subin Kim (MIT). This fork never falls back to browser cookies on its own: cookies are read only when you ask for them.

<a href="docs/assets/gh-attach-demo.webp">
  <img src="docs/assets/gh-attach-demo.webp" alt="gh-attach demo" width="800" />
</a>

</div>

## Quick Start

```sh
gh extension install Sygix/gh-attach
gh attach ./image.png -R owner/repo
```

## Installation

Requires [GitHub CLI](https://github.com/cli/cli#installation). Browser-cookie mode also requires `gh auth login`. Session-token mode does not when the repository ID is available from GitHub's page.

```sh
gh extension install Sygix/gh-attach
gh attach ./image.png -R owner/repo
```

## Usage

```sh
$ gh attach ./image.png -R owner/repo
https://github.com/user-attachments/assets/550e8400-e29b-41d4-a716-446655440000

$ gh attach upload ./image.png -R owner/repo
https://github.com/user-attachments/assets/550e8400-e29b-41d4-a716-446655440000

$ gh attach ./image.png ./report.pdf -R owner/repo
https://github.com/user-attachments/assets/550e8400-e29b-41d4-a716-446655440000
https://github.com/user-attachments/files/123/report.pdf

$ gh attach ./image.png ./report.pdf --markdown
![image.png](https://github.com/user-attachments/assets/550e8400-e29b-41d4-a716-446655440000)
[report.pdf](https://github.com/user-attachments/files/123/report.pdf)

$ gh attach ./image.png -R owner/repo --browser chrome --profile Default
https://github.com/user-attachments/assets/550e8400-e29b-41d4-a716-446655440000

$ GH_ATTACH_SESSION_TOKEN="$USER_SESSION" gh attach ./image.png -R owner/repo
https://github.com/user-attachments/assets/550e8400-e29b-41d4-a716-446655440000

$ gh attach ./image.png --json id,href,name
[
  {
    "href": "https://github.com/user-attachments/assets/550e8400-e29b-41d4-a716-446655440000",
    "id": 123,
    "name": "image.png"
  }
]

$ gh attach ./image.png --json href --jq '.[].href'
https://github.com/user-attachments/assets/550e8400-e29b-41d4-a716-446655440000

$ gh attach ./image.png --json href,name --template '{{range .}}{{.name}} -> {{.href}}{{"\n"}}{{end}}'
image.png -> https://github.com/user-attachments/assets/550e8400-e29b-41d4-a716-446655440000
```

Up to two files are uploaded concurrently, and per-file failures do not stop the remaining uploads. With `--json`, results are always returned as an array.

GitHub Cloud media uploads use the active `gh` token only; an authentication or permission error is returned as is, with no browser-cookie fallback. Supported formats are PNG, JPEG, GIF, WebP, SVG (up to 10 MiB), and MP4, MOV, WebM (up to 100 MiB; plan limits may be lower). GHES, other file types and JSON fields `id` or `originalName` need session cookies, which are read only when requested with `--browser`, `--profile`, `--cookie-store-path` or `--session-token` / `GH_ATTACH_SESSION_TOKEN`. Downloads never retry with browser cookies on their own either.

Embed uploaded URLs in an issue, PR, or comment and verify viewer access; unattached uploads may return 404 to others.

### Download

```sh
$ gh attach download https://github.com/user-attachments/assets/550e8400-e29b-41d4-a716-446655440000 -O image.png

$ gh attach download https://github.com/user-attachments/files/123/report.pdf -O -
```

Downloads use an explicit session token or browser selection first. Otherwise they use the active `gh` authentication token and retry with the matching browser cookies after an authorization failure, including GitHub sign-in or SSO pages returned with HTTP 200.

If browser authentication still returns a sign-in or SSO page, the download fails before writing the output. Sign in to GitHub and complete SSO in the selected browser, then retry.

## Use with AI agents

`gh-attach` ships as an [agent skill](https://agentskills.io), so AI coding agents can attach screenshots or files to a PR or issue, embed them as Markdown, and download `user-attachments` URLs from a natural-language request like *"attach this screenshot to the PR"*.

```sh
npx skills add Sygix/gh-attach
```

The [Agent Skills standard](https://agentskills.io/clients) is supported by Claude Code, OpenAI Codex, Cursor, GitHub Copilot, and more.

## Options

### Upload options

- `-R, --repo <[HOST/]OWNER/REPO>`: Target repository. Auto detection is available from current repository.
- `--browser <name>`: Browser to read cookies from (`auto|arc|atlas|brave|chrome|chromium|comet|dia|edge|firefox|floorp|helium|librewolf|opera|safari|vivaldi|waterfox|whale|zen`).
- `--profile <name>`: Browser profile name. For Firefox-family multi-account containers, append `:<container-name>` or `:id=<container-id>` to pin a specific container (e.g. `default:Work`, `default:id=2`).
- `--cookie-store-path <path>`: Explicit cookie DB file path.
- `--session-token <value>`: Bare `user_session` cookie value. Prefer the `GH_ATTACH_SESSION_TOKEN` environment variable to keep this account credential out of command history and process arguments. Explicit browser options override the environment variable.
- `--markdown`: Output Markdown references.
- `--json <fields>`: Output JSON with selected fields.
- `-q, --jq <expression>`: Apply jq filter to JSON output (requires `--json`).
- `-t, --template <go-template>`: Format JSON output using Go template (requires `--json`).
- `-v, --verbose`: Print authentication logs to stderr.
- `-h, --help`: Show help.

### Download options

- `-O, --output <file>`: File to write to. Use `-` for standard output.
- `--clobber`: Overwrite an existing file.
- `--browser`, `--profile`, `--cookie-store-path`, `--session-token`: Select the same authentication sources as upload.
- `-v, --verbose`: Print cookie source resolution logs to stderr.
- `-h, --help`: Show help.

## Configuration

You can use a config file to register frequently used browser settings without having to pass them as command line arguments each time.

The config file is loaded from `${XDG_CONFIG_HOME:-~/.config}/gh/attach.yml`.

**Example**

```yaml
browsers:
  - browser: chrome
    profile: Default
  - browser: firefox
    profile: default-release
  - browser: safari
```

**Schema**

- `browser`: Browser to read cookies from (Required, one of `auto|arc|atlas|brave|chrome|chromium|comet|dia|edge|firefox|floorp|helium|librewolf|opera|safari|vivaldi|waterfox|whale|zen`)
- `profile`: Browser profile name/path (Optional, name or path)
- `cookie_store_path`: Explicit cookie DB file path (Optional)

## Supported Browsers

- Chromium family (Arc, Atlas, Brave, Chrome, Chromium, Comet, Dia, Edge, Helium, Opera, Vivaldi, Whale)
- Firefox family (Firefox, Floorp, LibreWolf, Waterfox, Zen)
- Safari

## How It Works

- It first resolves the target repository (`owner/repo`).
- GitHub Cloud media uses `https://uploads.<host>/user-attachments/assets`; the steps below describe session-cookie uploads.
- With browser-cookie mode, it resolves the current GitHub login via the `gh` API and selects a browser session whose [`dotcom_user`](https://docs.github.com/en/site-policy/privacy-policies/github-cookies#cookies) matches it.
- With `--session-token` or `GH_ATTACH_SESSION_TOKEN`, it uses the supplied bare `user_session` value without reading a browser.
- Using that session cookie, it requests GitHub upload policies (`/upload/policies/assets`) and uploads each file binary.
- It finalizes each user-attachments asset and prints the results as URLs or formatted output via `--json`.
- Downloads prefer an explicit authentication source, then the active `gh` token, then matching browser cookies.

## Development

```sh
go test ./...
go build ./...
```

## License

MIT, see [LICENSE](./LICENSE).
