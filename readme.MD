# tree 

golang.tokyo #21 DevQuiz課題

# インストール方法

```
# make
# bin/tree testdata/case1
testdata/case1
├── dir1
│   ├── dir11
│   │   ├── file2
│   │   └── file3
│   └── dir12
│       ├── file4
│       └── file5
└── file1

3 directories, 5 files
```

#  使用方法

```
Usage: tree: bin/tree [OPTION] dir_pass
    -L deps
        階層の深さ。0(デフォルト)で無制限
    -D 
        ディレクトリのみ表示
    -G 
        ファイル名をGopherでマスク
```

