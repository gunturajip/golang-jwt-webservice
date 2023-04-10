export .env vars in local dev:
`
export $(grep -v '^#' .env | xargs)
`