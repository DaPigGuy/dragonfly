package block

import (
	_ "github.com/df-mc/dragonfly/server/internal/block_internal"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	_ "unsafe" // Imported for compiler directives.
)

//go:generate go run ../../cmd/blockhash -o hash.go .

// init registers all blocks implemented by Dragonfly.
func init() {
	world.RegisterBlock(Air{})
	world.RegisterBlock(Stone{})
	world.RegisterBlock(Stone{Smooth: true})
	world.RegisterBlock(Granite{})
	world.RegisterBlock(Granite{Polished: true})
	world.RegisterBlock(Diorite{})
	world.RegisterBlock(Diorite{Polished: true})
	world.RegisterBlock(Andesite{})
	world.RegisterBlock(Andesite{Polished: true})
	world.RegisterBlock(Grass{})
	world.RegisterBlock(DirtPath{})
	world.RegisterBlock(Dirt{})
	world.RegisterBlock(Dirt{Coarse: true})
	world.RegisterBlock(Cobblestone{})
	world.RegisterBlock(Cobblestone{Mossy: true})
	world.RegisterBlock(Bedrock{})
	world.RegisterBlock(Bedrock{InfiniteBurning: true})
	world.RegisterBlock(Obsidian{})
	world.RegisterBlock(Obsidian{Crying: true})
	world.RegisterBlock(DiamondBlock{})
	world.RegisterBlock(Glass{})
	world.RegisterBlock(Glowstone{})
	world.RegisterBlock(EmeraldBlock{})
	world.RegisterBlock(EndBricks{})
	world.RegisterBlock(GoldBlock{})
	world.RegisterBlock(NetheriteBlock{})
	world.RegisterBlock(IronBlock{})
	world.RegisterBlock(CoalBlock{})
	world.RegisterBlock(Beacon{})
	world.RegisterBlock(Sponge{})
	world.RegisterBlock(Sponge{Wet: true})
	world.RegisterBlock(LapisBlock{})
	world.RegisterBlock(Terracotta{})
	world.RegisterBlock(GlassPane{})
	world.RegisterBlock(IronBars{})
	world.RegisterBlock(NetherBrickFence{})
	world.RegisterBlock(EndStone{})
	world.RegisterBlock(Netherrack{})
	world.RegisterBlock(QuartzBricks{})
	world.RegisterBlock(Clay{})
	world.RegisterBlock(AncientDebris{})
	world.RegisterBlock(NetherGoldOre{})
	world.RegisterBlock(NetherQuartzOre{})
	world.RegisterBlock(Melon{})
	world.RegisterBlock(Sand{})
	world.RegisterBlock(Sand{Red: true})
	world.RegisterBlock(Gravel{})
	world.RegisterBlock(Bricks{})
	world.RegisterBlock(SoulSand{})
	world.RegisterBlock(Barrier{})
	world.RegisterBlock(SeaLantern{})
	world.RegisterBlock(SoulSoil{})
	world.RegisterBlock(BlueIce{})
	world.RegisterBlock(GildedBlackstone{})
	world.RegisterBlock(Shroomlight{})
	world.RegisterBlock(InvisibleBedrock{})
	world.RegisterBlock(DragonEgg{})
	world.RegisterBlock(NoteBlock{})
	world.RegisterBlock(NetherSprouts{})
	world.RegisterBlock(Tuff{})
	world.RegisterBlock(Calcite{})
	for _, ore := range OreTypes() {
		world.RegisterBlock(CoalOre{Type: ore})
		world.RegisterBlock(IronOre{Type: ore})
		world.RegisterBlock(GoldOre{Type: ore})
		world.RegisterBlock(CopperOre{Type: ore})
		world.RegisterBlock(LapisOre{Type: ore})
		world.RegisterBlock(DiamondOre{Type: ore})
		world.RegisterBlock(EmeraldOre{Type: ore})
	}
	world.RegisterBlock(RawIronBlock{})
	world.RegisterBlock(RawGoldBlock{})
	world.RegisterBlock(RawCopperBlock{})
	world.RegisterBlock(MossCarpet{})
	world.RegisterBlock(SporeBlossom{})
	world.RegisterBlock(Dripstone{})
	world.RegisterBlock(DriedKelpBlock{})
	world.RegisterBlock(HoneycombBlock{})
	world.RegisterBlock(Podzol{})

	registerAll(allBarrels())
	registerAll(allBasalt())
	registerAll(allBeetroot())
	registerAll(allBoneBlock())
	registerAll(allCake())
	registerAll(allCarpet())
	registerAll(allCarrots())
	registerAll(allChests())
	registerAll(allConcrete())
	registerAll(allConcretePowder())
	registerAll(allCocoaBeans())
	registerAll(allCoral())
	registerAll(allCoralBlocks())
	registerAll(allEndBrickStairs())
	registerAll(allWool())
	registerAll(allStainedTerracotta())
	registerAll(allGlazedTerracotta())
	registerAll(allStainedGlass())
	registerAll(allStainedGlassPane())
	registerAll(allLanterns())
	registerAll(allFire())
	registerAll(allPlanks())
	registerAll(allFence())
	registerAll(allFenceGates())
	registerAll(allWoodStairs())
	registerAll(allDoors())
	registerAll(allTrapdoors())
	registerAll(allWoodSlabs())
	registerAll(allLogs())
	registerAll(allLeaves())
	registerAll(allTorches())
	registerAll(allPumpkinStems())
	registerAll(allPumpkins())
	registerAll(allLitPumpkins())
	registerAll(allMelonStems())
	registerAll(allFarmland())
	registerAll(allLava())
	registerAll(allWater())
	registerAll(allKelp())
	registerAll(allPotato())
	registerAll(allWheat())
	registerAll(allQuartz())
	registerAll(allNetherWart())
	registerAll(allTallGrass())
	registerAll(allDoubleTallGrass())
	registerAll(allSandstones())
	registerAll(allDoubleFlowers())
	registerAll(allFlowers())
	registerAll(allPrismarine())
	registerAll(allLight())
}

func init() {
	world.RegisterItem(Air{})
	world.RegisterItem(Stone{})
	world.RegisterItem(Stone{Smooth: true})
	world.RegisterItem(Granite{})
	world.RegisterItem(Granite{Polished: true})
	world.RegisterItem(Diorite{})
	world.RegisterItem(Diorite{Polished: true})
	world.RegisterItem(Andesite{})
	world.RegisterItem(Andesite{Polished: true})
	world.RegisterItem(Grass{})
	world.RegisterItem(DirtPath{})
	world.RegisterItem(Dirt{})
	world.RegisterItem(Dirt{Coarse: true})
	world.RegisterItem(Cobblestone{})
	world.RegisterItem(Bedrock{})
	world.RegisterItem(Kelp{})
	world.RegisterItem(Log{Wood: OakWood()})
	world.RegisterItem(Log{Wood: SpruceWood()})
	world.RegisterItem(Log{Wood: BirchWood()})
	world.RegisterItem(Log{Wood: JungleWood()})
	world.RegisterItem(Leaves{Wood: OakWood(), Persistent: true})
	world.RegisterItem(Leaves{Wood: SpruceWood(), Persistent: true})
	world.RegisterItem(Leaves{Wood: BirchWood(), Persistent: true})
	world.RegisterItem(Leaves{Wood: JungleWood(), Persistent: true})
	world.RegisterItem(Chest{})
	world.RegisterItem(Cobblestone{Mossy: true})
	world.RegisterItem(Leaves{Wood: AcaciaWood(), Persistent: true})
	world.RegisterItem(Leaves{Wood: DarkOakWood(), Persistent: true})
	world.RegisterItem(Log{Wood: AcaciaWood()})
	world.RegisterItem(Log{Wood: DarkOakWood()})
	world.RegisterItem(Log{Wood: SpruceWood(), Stripped: true})
	world.RegisterItem(Log{Wood: BirchWood(), Stripped: true})
	world.RegisterItem(Log{Wood: JungleWood(), Stripped: true})
	world.RegisterItem(Log{Wood: AcaciaWood(), Stripped: true})
	world.RegisterItem(Log{Wood: DarkOakWood(), Stripped: true})
	world.RegisterItem(Log{Wood: OakWood(), Stripped: true})
	for _, c := range item.Colours() {
		world.RegisterItem(Concrete{Colour: c})
		world.RegisterItem(ConcretePowder{Colour: c})
		world.RegisterItem(StainedTerracotta{Colour: c})
		world.RegisterItem(Carpet{Colour: c})
		world.RegisterItem(Wool{Colour: c})
		world.RegisterItem(StainedGlass{Colour: c})
		world.RegisterItem(StainedGlassPane{Colour: c})
		world.RegisterItem(GlazedTerracotta{Colour: c})
	}
	for _, b := range allLight() {
		world.RegisterItem(b.(world.Item))
	}
	for _, w := range WoodTypes() {
		world.RegisterItem(Planks{Wood: w})
	}
	world.RegisterItem(WoodStairs{Wood: OakWood()})
	world.RegisterItem(WoodStairs{Wood: SpruceWood()})
	world.RegisterItem(WoodStairs{Wood: BirchWood()})
	world.RegisterItem(WoodStairs{Wood: JungleWood()})
	world.RegisterItem(WoodStairs{Wood: AcaciaWood()})
	world.RegisterItem(WoodStairs{Wood: DarkOakWood()})
	world.RegisterItem(WoodStairs{Wood: CrimsonWood()})
	world.RegisterItem(WoodStairs{Wood: WarpedWood()})
	world.RegisterItem(WoodSlab{Wood: OakWood()})
	world.RegisterItem(WoodSlab{Wood: SpruceWood()})
	world.RegisterItem(WoodSlab{Wood: BirchWood()})
	world.RegisterItem(WoodSlab{Wood: JungleWood()})
	world.RegisterItem(WoodSlab{Wood: AcaciaWood()})
	world.RegisterItem(WoodSlab{Wood: DarkOakWood()})
	world.RegisterItem(WoodSlab{Wood: CrimsonWood()})
	world.RegisterItem(WoodSlab{Wood: WarpedWood()})
	world.RegisterItem(WoodSlab{Wood: OakWood(), Double: true})
	world.RegisterItem(WoodSlab{Wood: SpruceWood(), Double: true})
	world.RegisterItem(WoodSlab{Wood: BirchWood(), Double: true})
	world.RegisterItem(WoodSlab{Wood: JungleWood(), Double: true})
	world.RegisterItem(WoodSlab{Wood: AcaciaWood(), Double: true})
	world.RegisterItem(WoodSlab{Wood: DarkOakWood(), Double: true})
	world.RegisterItem(WoodSlab{Wood: CrimsonWood(), Double: true})
	world.RegisterItem(WoodSlab{Wood: WarpedWood(), Double: true})
	world.RegisterItem(Obsidian{})
	world.RegisterItem(Obsidian{Crying: true})
	world.RegisterItem(DiamondBlock{})
	world.RegisterItem(Glass{})
	world.RegisterItem(Glowstone{})
	world.RegisterItem(EmeraldBlock{})
	world.RegisterItem(EndBricks{})
	world.RegisterItem(EndBrickStairs{})
	world.RegisterItem(NetheriteBlock{})
	world.RegisterItem(GoldBlock{})
	world.RegisterItem(IronBlock{})
	world.RegisterItem(CoalBlock{})
	world.RegisterItem(Beacon{})
	world.RegisterItem(Sponge{})
	world.RegisterItem(Sponge{Wet: true})
	world.RegisterItem(LapisBlock{})
	world.RegisterItem(Terracotta{})
	world.RegisterItem(Quartz{})
	world.RegisterItem(Quartz{Smooth: true})
	world.RegisterItem(ChiseledQuartz{})
	world.RegisterItem(QuartzPillar{})
	world.RegisterItem(QuartzBricks{})
	world.RegisterItem(GlassPane{})
	world.RegisterItem(IronBars{})
	world.RegisterItem(WoodFence{Wood: OakWood()})
	world.RegisterItem(WoodFence{Wood: SpruceWood()})
	world.RegisterItem(WoodFence{Wood: BirchWood()})
	world.RegisterItem(WoodFence{Wood: JungleWood()})
	world.RegisterItem(WoodFence{Wood: AcaciaWood()})
	world.RegisterItem(WoodFence{Wood: DarkOakWood()})
	world.RegisterItem(WoodFence{Wood: CrimsonWood()})
	world.RegisterItem(WoodFence{Wood: WarpedWood()})
	world.RegisterItem(NetherBrickFence{})
	world.RegisterItem(WoodFenceGate{Wood: OakWood()})
	world.RegisterItem(WoodFenceGate{Wood: SpruceWood()})
	world.RegisterItem(WoodFenceGate{Wood: BirchWood()})
	world.RegisterItem(WoodFenceGate{Wood: JungleWood()})
	world.RegisterItem(WoodFenceGate{Wood: AcaciaWood()})
	world.RegisterItem(WoodFenceGate{Wood: DarkOakWood()})
	world.RegisterItem(WoodFenceGate{Wood: CrimsonWood()})
	world.RegisterItem(WoodFenceGate{Wood: WarpedWood()})
	world.RegisterItem(WoodTrapdoor{Wood: OakWood()})
	world.RegisterItem(WoodTrapdoor{Wood: SpruceWood()})
	world.RegisterItem(WoodTrapdoor{Wood: BirchWood()})
	world.RegisterItem(WoodTrapdoor{Wood: JungleWood()})
	world.RegisterItem(WoodTrapdoor{Wood: AcaciaWood()})
	world.RegisterItem(WoodTrapdoor{Wood: DarkOakWood()})
	world.RegisterItem(WoodTrapdoor{Wood: CrimsonWood()})
	world.RegisterItem(WoodTrapdoor{Wood: WarpedWood()})
	world.RegisterItem(WoodDoor{Wood: OakWood()})
	world.RegisterItem(WoodDoor{Wood: SpruceWood()})
	world.RegisterItem(WoodDoor{Wood: BirchWood()})
	world.RegisterItem(WoodDoor{Wood: JungleWood()})
	world.RegisterItem(WoodDoor{Wood: AcaciaWood()})
	world.RegisterItem(WoodDoor{Wood: DarkOakWood()})
	world.RegisterItem(WoodDoor{Wood: CrimsonWood()})
	world.RegisterItem(WoodDoor{Wood: WarpedWood()})
	for _, c := range allCoral() {
		world.RegisterItem(c.(world.Item))
	}
	for _, c := range allCoralBlocks() {
		world.RegisterItem(c.(world.Item))
	}
	for _, s := range allSandstones() {
		world.RegisterItem(s.(world.Item))
	}
	world.RegisterItem(Barrel{})
	for _, ore := range OreTypes() {
		world.RegisterItem(CoalOre{Type: ore})
		world.RegisterItem(IronOre{Type: ore})
		world.RegisterItem(GoldOre{Type: ore})
		world.RegisterItem(CopperOre{Type: ore})
		world.RegisterItem(LapisOre{Type: ore})
		world.RegisterItem(DiamondOre{Type: ore})
		world.RegisterItem(EmeraldOre{Type: ore})
	}
	world.RegisterItem(Pumpkin{})
	world.RegisterItem(LitPumpkin{})
	world.RegisterItem(Pumpkin{Carved: true})
	world.RegisterItem(EndStone{})
	world.RegisterItem(Netherrack{})
	world.RegisterItem(Clay{})
	world.RegisterItem(BoneBlock{})
	world.RegisterItem(Lantern{Type: NormalFire()})
	world.RegisterItem(Lantern{Type: SoulFire()})
	world.RegisterItem(AncientDebris{})
	world.RegisterItem(NetherGoldOre{})
	world.RegisterItem(NetherQuartzOre{})
	world.RegisterItem(CocoaBean{})
	world.RegisterItem(WheatSeeds{})
	world.RegisterItem(BeetrootSeeds{})
	world.RegisterItem(Potato{})
	world.RegisterItem(Carrot{})
	world.RegisterItem(PumpkinSeeds{})
	world.RegisterItem(MelonSeeds{})
	world.RegisterItem(Melon{})
	world.RegisterItem(Sand{})
	world.RegisterItem(Sand{Red: true})
	world.RegisterItem(Gravel{})
	world.RegisterItem(Bricks{})
	world.RegisterItem(SoulSand{})
	world.RegisterItem(Barrier{})
	world.RegisterItem(Basalt{})
	world.RegisterItem(Basalt{Polished: true})
	world.RegisterItem(SeaLantern{})
	world.RegisterItem(SoulSoil{})
	world.RegisterItem(BlueIce{})
	world.RegisterItem(GildedBlackstone{})
	world.RegisterItem(Shroomlight{})
	world.RegisterItem(Torch{Type: NormalFire()})
	world.RegisterItem(Torch{Type: SoulFire()})
	world.RegisterItem(Cake{})
	world.RegisterItem(NetherWart{})
	world.RegisterItem(InvisibleBedrock{})
	world.RegisterItem(NoteBlock{Pitch: 24})
	world.RegisterItem(DragonEgg{})
	world.RegisterItem(TallGrass{})
	world.RegisterItem(TallGrass{Type: Fern()})
	world.RegisterItem(DoubleTallGrass{})
	world.RegisterItem(DoubleTallGrass{Type: Fern()})
	world.RegisterItem(DoubleFlower{Type: Sunflower()})
	world.RegisterItem(DoubleFlower{Type: Lilac()})
	world.RegisterItem(DoubleFlower{Type: RoseBush()})
	world.RegisterItem(DoubleFlower{Type: Peony()})
	world.RegisterItem(NetherSprouts{})
	world.RegisterItem(Farmland{})
	world.RegisterItem(Tuff{})
	world.RegisterItem(Calcite{})
	world.RegisterItem(RawIronBlock{})
	world.RegisterItem(RawGoldBlock{})
	world.RegisterItem(RawCopperBlock{})
	world.RegisterItem(MossCarpet{})
	world.RegisterItem(SporeBlossom{})
	world.RegisterItem(Dripstone{})
	world.RegisterItem(Flower{Type: Dandelion()})
	world.RegisterItem(Flower{Type: Poppy()})
	world.RegisterItem(Flower{Type: BlueOrchid()})
	world.RegisterItem(Flower{Type: Allium()})
	world.RegisterItem(Flower{Type: AzureBluet()})
	world.RegisterItem(Flower{Type: RedTulip()})
	world.RegisterItem(Flower{Type: OrangeTulip()})
	world.RegisterItem(Flower{Type: WhiteTulip()})
	world.RegisterItem(Flower{Type: PinkTulip()})
	world.RegisterItem(Flower{Type: OxeyeDaisy()})
	world.RegisterItem(Flower{Type: Cornflower()})
	world.RegisterItem(Flower{Type: LilyOfTheValley()})
	world.RegisterItem(Flower{Type: WitherRose()})
	world.RegisterItem(Prismarine{Type: NormalPrismarine()})
	world.RegisterItem(Prismarine{Type: DarkPrismarine()})
	world.RegisterItem(Prismarine{Type: BrickPrismarine()})
	world.RegisterItem(DriedKelpBlock{})
	world.RegisterItem(HoneycombBlock{})
	world.RegisterItem(Podzol{})

	world.RegisterItem(item.Bucket{Content: Water{}})
	world.RegisterItem(item.Bucket{Content: Lava{}})
}

// readSlice reads an interface slice from a map at the key passed.
//noinspection GoCommentLeadingSpace
func readSlice(m map[string]interface{}, key string) []interface{} {
	//lint:ignore S1005 Double assignment is done explicitly to prevent panics.
	v, _ := m[key]
	b, _ := v.([]interface{})
	return b
}

// readString reads a string from a map at the key passed.
//noinspection GoCommentLeadingSpace
func readString(m map[string]interface{}, key string) string {
	//lint:ignore S1005 Double assignment is done explicitly to prevent panics.
	v, _ := m[key]
	b, _ := v.(string)
	return b
}

// readInt32 reads an int32 from a map at the key passed.
//noinspection GoCommentLeadingSpace
func readInt32(m map[string]interface{}, key string) int32 {
	//lint:ignore S1005 Double assignment is done explicitly to prevent panics.
	v, _ := m[key]
	b, _ := v.(int32)
	return b
}

// readByte reads a byte from a map at the key passed.
//noinspection GoCommentLeadingSpace
func readByte(m map[string]interface{}, key string) byte {
	//lint:ignore S1005 Double assignment is done explicitly to prevent panics.
	v, _ := m[key]
	b, _ := v.(byte)
	return b
}

func registerAll(blocks []world.Block) {
	for _, b := range blocks {
		world.RegisterBlock(b)
	}
}
