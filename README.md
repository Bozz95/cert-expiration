# cert-expiration

Go tool to check how many days till the expiration of a website certificate

## Setup

1. Install Go (1.19 or later)
2. Clone this repository
3. Run the tool:

   ```bash
   go run main.go -url https://example.com
   ```

## Usage

```bash
go run main.go -url <https-url>
```

### Arguments

- `-url`: Https URL which certification expiration will be checked. Required

### Examples

```bash
# Check Google's certificate
go run main.go -url https://google.com

# Check GitHub's certificate
go run main.go -url https://github.com

# Test it on a custom port
go run main.go -url https://example.com:1234
```

## Docker Usage

### Build the image

```bash
# Build with SBOM
docker build --sbom=true -t cert-expiration .

# Or build without SBOM
docker build -t cert-expiration .
```

### Run the container

```bash
# Check Google's certificate
docker run --rm cert-expiration -url https://google.com

# Check GitHub's certificate
docker run --rm cert-expiration -url https://github.com

# Check any HTTPS site
docker run --rm cert-expiration -url https://example.com
```

## CI/CD

The project uses GitHub Actions for automated testing and releases:

- **On push to main**: Runs tests, generates semantic version tags, and publishes Docker images
- **Semantic versioning**: Based on conventional commits
  - `feat:` → minor version bump
  - `fix:` → patch version bump  
  - `BREAKING CHANGE:` → major version bump
- **Container images**: Published to `ghcr.io` with SBOM

### Using published images

```bash
# Use latest version
docker run --rm ghcr.io/your-username/cert-expiration:latest -url https://google.com

# Use specific version
docker run --rm ghcr.io/your-username/cert-expiration:v1.0.0 -url https://google.com
```

## Testing

Run the unit tests:

```bash
# Run all tests
go test ./test

# Run tests with verbose output
go test -v ./test

# Run tests with coverage
go test -cover ./test
```
