# horizon

> A private data hording environment

## Dependencies

- Run `get-deps.sh`

## todo

- [ ] [package migrations with bindata](https://github.com/rubenv/sql-migrate#embedding-migrations-with-bindata)
- [ ] Setup encrypted S3 as storage? Store as blocks of files? Any fuse-related go libraries possible to use?
- [ ] Setup psql ts_vectors for name search
  - [ ] then we can use https://www.census.gov/econ/cbp/download/ and the zip code files ([layout](https://www.census.gov/econ/cbp/download/noise_layout/ZIP_Totals_Layout10.txt))

#### todo (libraries)

- github.com/blevesearch/bleve
- github.com/koyachi/go-nude
- github.com/neurosnap/sentences
- github.com/soundcloud/roshi ?
