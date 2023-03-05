# magicedensniper
A sniper for Magic Eden that detects all listings, will expand usage to other marketplace in future. Autobuy is supported with API Key.

1. Using `.env.example` as example, create a new `.env` file with your `NODE_ENDPOINT`. `ME_APIKEY` and `PRIVATE_KEY` are needed for autobuy feature.

2. Configure your autobuy condition at `internal/sniper/sniper.go:120`

3. Run script with `go run cmd/main.go`

4. If run successfully, you will see logs as below at your standard output (usually terminal/cmd).
*Do note that the public Solana node has request limit so you will stop seeing results if you don't use your own node.
