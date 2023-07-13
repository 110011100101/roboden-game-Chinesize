package gamedata

// Season 0: build 11+
const SeasonNumber = 0

// Version 2
// - Added min movement radius to avoid accidental clicks
// - Fixed mismatching sprite size
//
// Version 3 (post jam)
// - Improved performance
// - Big map size option
// - New font
// - Updated drone cloning algorithm
// - Fixed a cloning price bug in colony planner
// - Fixed colony movement overjump
// - Fixed some typos
// - Fixed some text rendering issues (it was too blurry at times)
// - Made macos/osx build possible
// - Improved tutorial texts a bit
// - Improved colony planner & core code (agent search code)
// - Colonies now prefer to use green cloners
//
// # Version 4
//
// * Misc:
//   - Added Russian language support
//
// * New features:
//   - Added walls (in forms of land cracks)
//   - Added mountains as a higher height wall types
//   - Crawler creeps (spawned by a boss)
//   - Added mortar drone
//   - Added anti-air missile drone
//   - Added prism drone
//   - Added turrets
//
// * Visual improvements:
//   - Added a flash effect when a unit (friendly or not) takes damage
//
// * Gameplay:
//   - Unit cost changes (rebalancing)
//   - Creep bases are now guarded by a tower (usually)
//   - Higher resource priority decreases the upkeep (bonus caps at 80% priority with 60% reduction)
//   - Reworked most merge recipes
//   - Bases try to send at least 1 worker as reinforcements
//   - Base will try to produce at least 2 workers even if security priority is high
//   - Rework faction passive bonuses
//
// * Fixes:
//   - Drones no longer try to pick up a depleted resource
//   - Fixed an upkeed visual bug (trash is rendered near the base)
//   - Fixed invalid drone stats (max hp and speed)
//
// # Version 5
//
// * New features:
//   - Experimental mobile devices support
//   - Finished controllers (gamepad devices) support
//   - Crawlers now know a "scatter" behavior
//
// * UX:
//   - The colony selection is more precise now
//   - Added "toggle base" and "menu" (burger) buttons
//
// * Visual improvements:
//   - Added icons for the 5th option
//   - Reworked sprites for the action options
//
// * Fixes:
//   - Drones with charging mode no longer confuse creeps (we're clearing the waypoint now)
//   - Fixed crawlers "never scout" bug
//
// # Version 7 (was uploaded as version 6 by accident)
//
// * New features:
//   - Red crystals resource
//   - Drone ranks (normal, elite, super elite)
//
// * Gameplay
//   - Increased a colony drone limit
//   - Added more colony drone traits
//
// * UX:
//   - Made it clear which option was selected
//
// * Visual improvements:
//   - New action cooldown effect
//
// * Fixes:
//   - Upon defeat, hide menu and toggle buttons
//   - Fixed resource collection bug (drone cargo value stacking)
//
// * Performance:
//   - Use a pre-decoded ogg stream instead of decoding it on the fly
//   - Since 99% graphic objects are ge.Sprite, they're now stored as separate slices (less iface calls)
//
// # Version 8
// So many things happened, but I forgot to write them down.
//
// # Version 11
// * Added online leaderboard to the game
// * Added a separate rewards screen
// * Added teleports
// * Balance tweaks
// * Resources priority effect rework
// * Changed the default keyboard binds for camera to WASD
//
// # Version 12
//
// * New content:
//   - Added Scarab tier 2 drone
//   - Added Devourer tier 3 drone
//   - Added "Oil regeneration rate" world option
//   - Added "Terrain" world option
//   - Super creep versions (classic + arena)
//   - Added Seeker tier 2 drone
//
// * Balance:
//   - Add 2 max hp to Prism drones (28 -> 30)
//   - Make tether towers affect up to 4 drones (it was 2 before)
//   - Make tether towers slightly reduce the amount of drone energy consumption
//   - Made attack action range scale better
//   - Reworked (nerfed) couriers income
//
// * UX:
//   - Added difficulty description tag in the game lobby
//   - Added a UI toggle key
//   - Added a Timer extra option
//
// * Bug fixes:
//   - Fixed overlapping teleports bug
//   - Fixed universal drones not being queries by SearchWorkers in drones container
//   - Make "catch em all" achievement come with Elite grade right away
//   - Fixed multi rewards
//   - Fixed Trucker drone diode location
//   - Fixed leaderboard layout issues
//   - Removed "colony under attack" notice during the fatal damage
//   - Fixed some of the faction bonuses (they were not applied properly)
//
// * Optimizations:
//   - The game doesn't create animation objects for drones with no animation anymore
//   - Clusters for units (makes target selection faster)
//   - Less allocations in simulation mode
//   - Terrain is drawn as a single texture instead of a set of sprites
//
// * Visuals:
//   - Generate less looped land cracks.
//
// # Version 13 (pre-Steam itchio early access release)
//
// * New content:
//   - A completely new tutorial mode (the intro mission)
//   - Added Commander tier 2 drone
//   - Added Targeter tier 2 drone
//   - Added Firebug tier 2 drone
//   - Added Harvester turret
//
// * Difficulty settings:
//   - Up to 5 creep bases in Classic mode (the previous limit was 4)
//   - Enable arena progression setting in infinite arena mode
//   - Change arena-related settings scaling (20% -> 25%)
//
// * Balance:
//   - Increased the max upkeep resources cost
//   - Freighter: now has a zero upkeep
//   - Roomba: increase upkeep (3 -> 9)
//
// * Creeps:
//   - Added a 10-20 sec delay before a howitzer can start firing its artillery
//
// * UX:
//   - Fixed camera snapping issue near the borders of the map
//   - The message window is not half-transparent
//   - Better gamepad configuration screen
//   - Better drones overview (with stats)
//   - Added seed option help on-hover text
//   - A notification for the colony being destroyed
//   - Added PlayStation and Nintendo Switch gamepads layout support
//   - Auto-pause when gamepad is disconnected
//   - In-game on-hover hints
//   - Hide the less relevant cursor when in the game
//
// * Visual improvements:
//   - Prism drone attack now has a hit impact effect
//   - Better layers for colonies (especially when they're landing)
//   - Better layers for explosions (they're above dreadnought now, as they should be)
//   - Better layout in some of the menus
//
// * Gameplay:
//   - Reworked the tutorial
//
// * Bug fixes:
//   - "Colony" settings tab is not properly disabled in Reverse mode
//   - Splash screen no longer depends on the classic mode fog of war settings
//   - No achievements in coop mode (this would be too easy)
//   - Fixed uncentered camera in a single bot mode
//   - Fixed a Seeker drone that was unlocked right away (should be unlocked at 1000 pts)
//   - Can no longer build on top of the teleporter
//
// * Performance:
//   - Precompute waypoint direction once for colony agents
//
// * Misc:
//   - Updated ebitenui library version
//   - Underscore character '_' is now allowed in the username
//
// # Version 14 (Steam release)
//
// * New content:
//   - A new secret achievement
const BuildNumber int = 14
