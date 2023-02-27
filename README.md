# SkinResizer

SkinResizer transforme vos skins au format 288x128 (32x32 par frame au lieu de 24x32)

## Installation

Téléchargez directement la dernière [release](https://github.com/HLerman/SkinResizer/releases/latest)

## Build

Si vous voulez build vous même le projet il sera nécessaire d'avoir [go](https://go.dev/doc/install) installé.

```bash
git clone https://github.com/HLerman/SkinResizer.git
cd SkinResizer
go build
```

## Usage

### Linux
```bash
./SkinResizer /home/user/slayersonline/Chipset/zeronin.png
```

ou pour transformer toutes les skins d'un repertoire :

```bash
find "/home/user/slayersonline/Chipset" -iname '*.png' -exec ./SkinResizer {} \;
```

### Windows
```powershell
./SkinResizer.exe 'C:\Users\User\Documents\Slayers Online\Chipset\zeronin.png'
```

ou pour transformer toutes les skins d'un repertoire :

```powershell
Get-ChildItem -Path 'C:\Users\User\Documents\Slayers Online\Chipset\*.png' | Foreach {./SkinResizer.exe $_.fullname}
```

## Informations

La skin nouvellement créée :
- n'est pas indexée
- a un fond transparent (ex: le fond rose de frost.png devient transparent)