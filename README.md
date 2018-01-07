# sticky-study-converter

Converts a specific rip of Minna No Nihongo transcription into Sticky Study Flash Cards format.

When run remember to remove some garbage character that occurs at the beginning of the file or it will come up with a wingding for some reason. was too lazy to check into it as this is a personal project anyways and the removal is trivial time spent.

## Exmaple

```shell
go run .\main.go -input .\.test\Minna_no_nihongo_1.11.txt -output .\.test\l11.csv -new
go run .\main.go -input .\.test\l11.csv -output .\.test\l11.csv.ou
```
