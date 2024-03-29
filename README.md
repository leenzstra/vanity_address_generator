# vanity_address_generator

## Vanity address generator for multiple address types [SOL, EVM, MOVE]

- Parallel address grinding
- Case (in)sensetive
- Selectable position in address
- Writes to file

## Usage

####  1. Source

```
Usage: go run cmd\gen.go [options]

  -c, --caseSensetive       is case sensetive (default false)
  -g, --generateCount int   address count to generate (default 1)
  -f, --logFile string      log file (default "logs")
  -s, --seq string          sequence to find (default "")
  -p, --seqPos int          start position of sequence in address (default 0)
  -t, --type string         address type (evm, solana, move)
  -w, --workers int         parallel workers count (default 10)
```

#### 2. Executable

Download executables for your OS from 'Releases' tab
```
    vanitygen [options]
```

## Examples
Generate 1 solana address
```
vanitygen -t sol

---
Output in logs file:
2024/03/15 18:59:06 start working
2024/03/15 18:59:06 pub: 3zB...VFn pk: 5BK...fYC
---
```
Generate 2 solana addresses starting with 'toly' case insensetive using 32 workers
```
vanitygen -t sol -w 32 -s toly -g 2

---
Output in logs file:
2024/03/15 19:09:43 start working
2024/03/15 19:09:44 pub: ToLyie...8Fb pk: 8o1...pbs
2024/03/15 19:10:23 pub: ToLyGg...hJL pk: 4RS...aQ4
---
```
Generate 1 solana address starting with 'wif' from 5 position case insensetive using 8 workers
```
vanitygen -t sol -w 8 -s wif -g 1 -p 5

---
Output in logs file:
2024/03/15 19:13:47 start working
2024/03/15 19:13:47 pub: 4kWBcwiFnBd...x61 pk: 3A6...Q8M
---
```
Generate 1 eth (evm) address starting with 'afa' (0x excluded).

❗ Only numbers 0-9 and letters a-f supported ❗
```
vanitygen -t evm -s afa 

---
Output in logs file:
2024/03/22 13:04:35 start working
2024/03/22 13:04:43 pub: 0xafa72...79c1b pk: 0x9d63...749b
---
```