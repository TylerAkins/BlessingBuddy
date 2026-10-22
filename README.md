<p align="center"><img src="media/logo.png" width="160" alt="BlessingBuddy logo"></p>

# BlessingBuddy

One-click buffs and rescue heals for friendly players and pets in World of Warcraft: Forever.

Target a friendly player (or pet) and BlessingBuddy shows a small bar with your single-target buffs, a recommended buff for the target's class, the target's health and a row of rescue heals. By default it appears for players **outside your group** only; party, raid, and self are optional in Interface > AddOns > BlessingBuddy.

For Paladins, Blessing of Salvation is a separate button for party and raid players when those visibility options are enabled.

## About this fork

BlessingBuddy was created by [Tom Steglitz](https://github.com/TomSteglitz/BlessingBuddy). This repository is a personal fork, and the original work remains credited to Tom Steglitz under the MIT license (see [LICENSE](LICENSE)). I made the changes in this fork for my own use. This fork adds optional visibility settings and group-only Blessing of Salvation.

I do not plan to distribute this fork on CurseForge unless the original addon is confirmed abandoned. The current automated release workflow still uploads tagged releases to CurseForge, so the workflow must be changed before that intention is reflected in future releases. The CurseForge link below points to the original addon listing; GitHub Releases are for this fork.

<p align="center">
  <img src="media/gallery-1-recommended-buff.png" width="640" alt="Kings already active, Might recommended next">
  <img src="media/gallery-2-out-of-range.png" width="640" alt="Icons turn red when the target is out of range">
  <img src="media/gallery-3-every-class.png" width="640" alt="Works for every buffing class, here a level 2 priest">
  <img src="media/gallery-4-pvp-warning.png" width="640" alt="Red PvP warning when helping the target would flag you">
  <img src="media/gallery-5-spellcasters.png" width="640" alt="A mage sees Arcane Intellect recommended for a priest">
</p>

## Download
- [CurseForge listing for the original addon](https://www.curseforge.com/wow/addons/blessingbuddy)
- [GitHub Releases for this fork](../../releases)

## Features
- Recommended buff by target class, skips buffs the target already has
- Paladin one-blessing rule aware (refreshes instead of overwriting)
- Rescue row: heals, shields, dispels + target health bar
- Heal prediction on the health bar: incoming heals, remaining healing of running HoTs, and a preview when you hover a heal button – overhealing is cut off at 100 %
- Paladin: your own blessing shows as done (greyed out) and turns into a refresh button with countdown when it has less than 5 minutes left
- Buff ranks match the target's level (no "target too low" errors)
- Pets of other players: health bar, heals, Thorns
- Works in combat, keybinds, range/cooldown display, PvP warning
- Always casts your highest learned rank
- Optional visibility for players outside your group, party members, raid members, and yourself (Interface > AddOns > BlessingBuddy; defaults keep the classic “strangers only” behavior outside combat; self off by default)
- English and German interface

Supported classes: Paladin, Druid, Priest, Mage, Warlock, Shaman.

## Commands
| Command | Effect |
|---|---|
| `/bb move` | show the bar to move it (Shift + drag title) |
| `/bb reset` | reset position |
| `/bb debug` | diagnostic info about your target |
| `/bb log on\|off\|clear` | write a detailed debug log to the SavedVariables file |
| `/bb healtest` | show which heal-prediction data the client provides |
| *(settings)* | Interface > AddOns > BlessingBuddy — who to show the bar for outside combat |

German aliases: `/segen move`, `/segen reset`, `/segen debug`.

## Reporting bugs
Please open an [issue](../../issues) and include your WoW version, class/level, the target and the output of `/bb debug`. For tricky bugs, `/bb log on`, reproduce, `/reload` and attach `WTF\Account\<ACCOUNT>\SavedVariables\BlessingBuddy.lua`.

## Releasing (maintainers)
1. Update `VERSION`, `BlessingBuddy.toc`, `CHANGELOG.md`, and `RELEASE_NOTES.md`. Versions use `MAJOR.MINOR.PATCH`, and the first heading in `RELEASE_NOTES.md` must match `VERSION`.
2. Commit the changes, then create and push the matching tag, for example: `git tag v1.6.7 && git push origin v1.6.7`.
3. The release workflow validates the version and notes, builds the ZIP, uploads it with the [BigWigs packager](https://github.com/BigWigsMods/packager), and creates a GitHub Release with the ZIP and manifest.

The current workflow also uploads to CurseForge when the repository has a `CF_API_KEY` secret. Remove or disable that upload before making a release under the distribution policy above.

## License
MIT – see [LICENSE](LICENSE).
