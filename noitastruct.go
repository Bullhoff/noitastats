package main

import ()

type EntityContent struct  {
	Name string `xml:"name,attr"`	// inventory_quick
	Tags string `xml:"tags,attr"`
	Transform struct  {
		PositionX string `xml:"position.x,attr"`
		PositionY string `xml:"position.y,attr"`
	} `xml:"_Transform"`
	AbilityComponent struct  {
		Gun_level string `xml:"gun_level,attr"`
		Ui_name string `xml:"ui_name.y,attr"`
		Mana string `xml:"mana.y,attr"`
		Mana_charge_speed string `xml:"mana_charge_speed.y,attr"`
		Mana_max string `xml:"mana_max.y,attr"`
		Reload_time_frames string `xml:"reload_time_frames.y,attr"`
		Rotate_hand_amount string `xml:"rotate_hand_amount.y,attr"`
	} `xml:"AbilityComponent"`
	ItemComponent struct  {
		Spawn_posX string `xml:"spawn_pos.x,attr"`
		Spawn_posY string `xml:"spawn_pos.y,attr"`
	} `xml:"ItemComponent"`
	ItemActionComponent struct  {
		Tags string `xml:"action_id,attr"`	// TELEPORT_PROJECTILE_SHORT
	}
	PhysicsImageShapeComponent struct  {
		Image_file string `xml:"image_file,attr"`	
		Material string `xml:"material,attr"`
	}
	ProjectileComponent struct  {
		Bounce_always string `xml:"bounce_always,attr"`	
		Bounce_at_any_angle string `xml:"bounce_at_any_angle,attr"`	
		Bounce_energy string `xml:"bounce_energy,attr"`	
		Camera_shake_when_shot string `xml:"camera_shake_when_shot,attr"`	
		Collect_materials_to_shooter string `xml:"collect_materials_to_shooter,attr"`	
		Collide_with_entities string `xml:"collide_with_entities,attr"`	
		Collide_with_tag string `xml:"collide_with_tag,attr"`	
		Collide_with_world string `xml:"collide_with_world,attr"`	
		Damage string `xml:"damage,attr"`	
		Damage_every_x_frames string `xml:"damage_every_x_frames,attr"`	
		Explosion_dont_damage_shooter string `xml:"explosion_dont_damage_shooter,attr"`	
		Friction string `xml:"friction,attr"`	
		Friendly_fire string `xml:"friendly_fire,attr"`	
		Lifetime string `xml:"lifetime,attr"`	
		Lifetime_randomness string `xml:"lifetime_randomness,attr"`	
		On_collision_die string `xml:"on_collision_die,attr"`	
		On_collision_remove_projectile string `xml:"on_collision_remove_projectile,attr"`	
		On_collision_spawn_entity string `xml:"on_collision_spawn_entity,attr"`	
		Speed_max string `xml:"speed_max,attr"`	
		Speed_min string `xml:"speed_min,attr"`	
		Config_explosion struct  {
			Damage string `xml:"damage,attr"`	
			Damage_mortals string `xml:"damage_mortals,attr"`
			Explosion_radius string `xml:"explosion_radius,attr"`
			Destroy_non_platform_solid_enabled string `xml:"destroy_non_platform_solid_enabled,attr"`
			Camera_shake string `xml:"camera_shake,attr"`
			
			Damage_critical struct  {
				Chance string `xml:"chance,attr"`	
				Damage_multiplier string `xml:"damage_multiplier,attr"`
			} `xml:"damage_critical"`
		} `xml:"config_explosion"`

		Damage_by_type struct  {
			Curse string `xml:"curse,attr"`	
			Drill string `xml:"drill,attr"`	
			Electricity string `xml:"electricity,attr"`	
			Explosion string `xml:"explosion,attr"`	
			Fire string `xml:"fire,attr"`	
			Healing string `xml:"healing,attr"`	
			Ice string `xml:"ice,attr"`	
			Melee string `xml:"melee,attr"`	
			Overeating string `xml:"overeating,attr"`	
			Physics_hit string `xml:"physics_hit,attr"`	
			Poison string `xml:"poison,attr"`	
			Projectile string `xml:"projectile,attr"`	
			Radioactive string `xml:"radioactive,attr"`	
			Slice string `xml:"slice,attr"`	
		} `xml:"damage_by_type"`
		Damage_critical struct  {
			Chance string `xml:"chance,attr"`	
			Damage_multiplier string `xml:"damage_multiplier,attr"`	
		} `xml:"damage_critical"`
		
	}
	VelocityComponent struct  {
		Affect_physics_bodies string `xml:"affect_physics_bodies,attr"`	
		Displace_liquid string `xml:"displace_liquid,attr"`	
		Gravity_x string `xml:"gravity_x,attr"`	
		Gravity_y string `xml:"gravity_y,attr"`	
		Mass string `xml:"mass,attr"`	
		Terminal_velocity string `xml:"terminal_velocity,attr"`	
	}
	
}

type WorldState struct {
	WorldStateComponent struct  {
		Lua_globals struct {
			E []struct {
				Key string `xml:"key,attr"`
				Value string `xml:"value,attr"`
			} `xml:"E"`
		} `xml:"lua_globals"`
	} `xml:"WorldStateComponent"`
}
type Stats struct {
	//[]Stats struct  {
		Filename string
		Deaths string `xml:"deaths,attr"`
		Kills string `xml:"kills,attr"`
		Player_kills string `xml:"player_kills,attr"`
		Player_projectile_count string `xml:"player_projectile_count,attr"`

		Death_map struct  {
			E struct  {
				Key string `xml:"key,attr"`
				Value string `xml:"value,attr"`
			}
		} `xml:"death_map"`

		Kill_map struct  {
			E struct  {
				Key string `xml:"key,attr"`
				Value string `xml:"value,attr"`
			}
		} `xml:"kill_map"`
	//}
}
type Sessions struct{
	Gameid string
	BUILD_NAME string `xml:"BUILD_NAME,attr"`
	Stats struct  {
		World_seed string `xml:"world_seed,attr"`
		Value string `xml:"value,attr"`
	} `xml:"stats"`
	Biome_baseline struct  {} `xml:"biome_baseline"`
	Item_map struct  {} `xml:"item_map"`
	Biomes_visited struct  {} `xml:"biomes_visited"`
	Deaths string `xml:"deaths,attr"`
	Kills string `xml:"kills,attr"`
	Player_kills string `xml:"player_kills,attr"`
	Player_projectile_count string `xml:"player_projectile_count,attr"`
	Death_map struct  {
		E struct  {
			Key string `xml:"key,attr"`
			Value string `xml:"value,attr"`
		}
	} `xml:"death_map"`
	Kill_map struct  {
		E struct  {
			Key string `xml:"key,attr"`
			Value string `xml:"value,attr"`
		}
	} `xml:"kill_map"`
}
type GunActionStruct struct {
	Description string `xml:",omitempty" json:"description,omitempty"`
	Lifetime_add string `xml:",omitempty" json:"lifetime_add,omitempty"`
	Id string `xml:",omitempty" json:"id,omitempty"`
	Mana string `xml:",omitempty" json:"mana,omitempty"`
	Name string `xml:",omitempty" json:"name,omitempty"`
	Price string `xml:",omitempty" json:"price,omitempty"`
	Related_extra_entities []string `xml:",omitempty" json:"related_extra_entities,omitempty"`
	Custom_xml_file []string `xml:",omitempty" json:"custom_xml_file,omitempty"`
	Related_projectiles []string `xml:",omitempty" json:"related_projectiles,omitempty"`
	Spawn_level string `xml:",omitempty" json:"spawn_level,omitempty"`
	Spawn_probability string `xml:",omitempty" json:"spawn_probability,omitempty"`
	Sprite string `xml:",omitempty" json:"sprite,omitempty"`
	Sprite_unidentified string `xml:",omitempty" json:"sprite_unidentified,omitempty"`
	Type string `xml:",omitempty" json:"type,omitempty"`

}

type ProjectileComponent struct {
	Speed_min string `xml:"speed_min,attr" json:"speed_min"`
	Speed_max string `xml:"speed_max,attr" json:"speed_max"`
	On_death_explode bool `xml:"on_death_explode,attr" json:"on_death_explode"`
	On_death_gfx_leave_sprite bool `xml:"on_death_gfx_leave_sprite,attr" json:"on_death_gfx_leave_sprite"`
	On_lifetime_out_explode bool `xml:"on_lifetime_out_explode,attr" json:"on_lifetime_out_explode"`
	On_collision_die bool `xml:"on_collision_die,attr" json:"on_collision_die"`
	Die_on_low_velocity bool `xml:"die_on_low_velocity,attr" json:"die_on_low_velocity"`
	Damage float32 `xml:"damage,attr" json:"damage"`
	Lifetime int `xml:"lifetime,attr" json:"lifetime"`
	Shoot_light_flash_r string `xml:"shoot_light_flash_r,attr" json:"shoot_light_flash_r"`
	Shoot_light_flash_g string `xml:"shoot_light_flash_g,attr" json:"shoot_light_flash_g"`
	Shoot_light_flash_b string `xml:"shoot_light_flash_b,attr" json:"shoot_light_flash_b"`
	Shoot_light_flash_radius string `xml:"shoot_light_flash_radius,attr" json:"shoot_light_flash_radius"`
	Knockback_force float32 `xml:"knockback_force,attr" json:"knockback_force"`
} 

type EntityStruct struct {
	ProjectileComponent ProjectileComponent `xml:"ProjectileComponent"`
	GunActions GunActionStruct `json:"GunActions"`
}

type Entity_Struct struct {
	GunActions []GunActionStruct `xml:",omitempty" json:"GunActions,omitempty"` 
	Tags string  `xml:"tags,attr,omitempty" json:"tags,omitempty"` 
	Name string  `xml:"name,attr,omitempty" json:"name,omitempty"` 
	Base []Base_Struct  `xml:"Base,omitempty" json:"Base,omitempty"` 
	ProjectileComponent []ProjectileComponent_Struct  `xml:"ProjectileComponent,omitempty" json:"ProjectileComponent,omitempty"` 
	SpriteComponent []SpriteComponent_Struct  `xml:"SpriteComponent,omitempty" json:"SpriteComponent,omitempty"` 
	ParticleEmitterComponent []ParticleEmitterComponent_Struct  `xml:"ParticleEmitterComponent,omitempty" json:"ParticleEmitterComponent,omitempty"` 
	HitboxComponent []HitboxComponent_Struct  `xml:"HitboxComponent,omitempty" json:"HitboxComponent,omitempty"` 
	DamageModelComponent []DamageModelComponent_Struct  `xml:"DamageModelComponent,omitempty" json:"DamageModelComponent,omitempty"` 
	AudioComponent []AudioComponent_Struct  `xml:"AudioComponent,omitempty" json:"AudioComponent,omitempty"` 
	VariableStorageComponent []VariableStorageComponent_Struct  `xml:"VariableStorageComponent,omitempty" json:"VariableStorageComponent,omitempty"` 
	Version string  `xml:"_version,attr,omitempty" json:"_version,omitempty"` 
	Serialize string  `xml:"serialize,attr,omitempty" json:"serialize,omitempty"` 
	Transform []Transform_Struct  `xml:"_Transform,omitempty" json:"_Transform,omitempty"` 
	AudioListenerComponent []AudioListenerComponent_Struct  `xml:"AudioListenerComponent,omitempty" json:"AudioListenerComponent,omitempty"` 
	AudioLoopComponent []AudioLoopComponent_Struct  `xml:"AudioLoopComponent,omitempty" json:"AudioLoopComponent,omitempty"` 
	CharacterDataComponent []CharacterDataComponent_Struct  `xml:"CharacterDataComponent,omitempty" json:"CharacterDataComponent,omitempty"` 
	CharacterPlatformingComponent []CharacterPlatformingComponent_Struct  `xml:"CharacterPlatformingComponent,omitempty" json:"CharacterPlatformingComponent,omitempty"` 
	ControlsComponent []ControlsComponent_Struct  `xml:"ControlsComponent,omitempty" json:"ControlsComponent,omitempty"` 
	DrugEffectComponent []DrugEffectComponent_Struct  `xml:"DrugEffectComponent,omitempty" json:"DrugEffectComponent,omitempty"` 
	GameLogComponent []GameLogComponent_Struct  `xml:"GameLogComponent,omitempty" json:"GameLogComponent,omitempty"` 
	GameStatsComponent []GameStatsComponent_Struct  `xml:"GameStatsComponent,omitempty" json:"GameStatsComponent,omitempty"` 
	GenomeDataComponent []GenomeDataComponent_Struct  `xml:"GenomeDataComponent,omitempty" json:"GenomeDataComponent,omitempty"` 
	GunComponent []GunComponent_Struct  `xml:"GunComponent,omitempty" json:"GunComponent,omitempty"` 
	HotspotComponent []HotspotComponent_Struct  `xml:"HotspotComponent,omitempty" json:"HotspotComponent,omitempty"` 
	IngestionComponent []IngestionComponent_Struct  `xml:"IngestionComponent,omitempty" json:"IngestionComponent,omitempty"` 
	Inventory2Component []Inventory2Component_Struct  `xml:"Inventory2Component,omitempty" json:"Inventory2Component,omitempty"` 
	InventoryGuiComponent []InventoryGuiComponent_Struct  `xml:"InventoryGuiComponent,omitempty" json:"InventoryGuiComponent,omitempty"` 
	ItemPickUpperComponent []ItemPickUpperComponent_Struct  `xml:"ItemPickUpperComponent,omitempty" json:"ItemPickUpperComponent,omitempty"` 
	KickComponent []KickComponent_Struct  `xml:"KickComponent,omitempty" json:"KickComponent,omitempty"` 
	LightComponent []LightComponent_Struct  `xml:"LightComponent,omitempty" json:"LightComponent,omitempty"` 
	LiquidDisplacerComponent []LiquidDisplacerComponent_Struct  `xml:"LiquidDisplacerComponent,omitempty" json:"LiquidDisplacerComponent,omitempty"` 
	LuaComponent []LuaComponent_Struct  `xml:"LuaComponent,omitempty" json:"LuaComponent,omitempty"` 
	MaterialInventoryComponent []MaterialInventoryComponent_Struct  `xml:"MaterialInventoryComponent,omitempty" json:"MaterialInventoryComponent,omitempty"` 
	MaterialSuckerComponent []MaterialSuckerComponent_Struct  `xml:"MaterialSuckerComponent,omitempty" json:"MaterialSuckerComponent,omitempty"` 
	PathFindingGridMarkerComponent []PathFindingGridMarkerComponent_Struct  `xml:"PathFindingGridMarkerComponent,omitempty" json:"PathFindingGridMarkerComponent,omitempty"` 
	PhysicsPickUpComponent []PhysicsPickUpComponent_Struct  `xml:"PhysicsPickUpComponent,omitempty" json:"PhysicsPickUpComponent,omitempty"` 
	PlatformShooterPlayerComponent []PlatformShooterPlayerComponent_Struct  `xml:"PlatformShooterPlayerComponent,omitempty" json:"PlatformShooterPlayerComponent,omitempty"` 
	PlayerCollisionComponent []PlayerCollisionComponent_Struct  `xml:"PlayerCollisionComponent,omitempty" json:"PlayerCollisionComponent,omitempty"` 
	ShotEffectComponent []ShotEffectComponent_Struct  `xml:"ShotEffectComponent,omitempty" json:"ShotEffectComponent,omitempty"` 
	SpriteAnimatorComponent []SpriteAnimatorComponent_Struct  `xml:"SpriteAnimatorComponent,omitempty" json:"SpriteAnimatorComponent,omitempty"` 
	SpriteParticleEmitterComponent []SpriteParticleEmitterComponent_Struct  `xml:"SpriteParticleEmitterComponent,omitempty" json:"SpriteParticleEmitterComponent,omitempty"` 
	SpriteStainsComponent []SpriteStainsComponent_Struct  `xml:"SpriteStainsComponent,omitempty" json:"SpriteStainsComponent,omitempty"` 
	StatusEffectDataComponent []StatusEffectDataComponent_Struct  `xml:"StatusEffectDataComponent,omitempty" json:"StatusEffectDataComponent,omitempty"` 
	StreamingKeepAliveComponent []StreamingKeepAliveComponent_Struct  `xml:"StreamingKeepAliveComponent,omitempty" json:"StreamingKeepAliveComponent,omitempty"` 
	TelekinesisComponent []TelekinesisComponent_Struct  `xml:"TelekinesisComponent,omitempty" json:"TelekinesisComponent,omitempty"` 
	VelocityComponent []VelocityComponent_Struct  `xml:"VelocityComponent,omitempty" json:"VelocityComponent,omitempty"` 
	WalletComponent []WalletComponent_Struct  `xml:"WalletComponent,omitempty" json:"WalletComponent,omitempty"` 
	Entity []Entity_Struct  `xml:"Entity,omitempty" json:"Entity,omitempty"` 
	InheritTransformComponent []InheritTransformComponent_Struct  `xml:"InheritTransformComponent,omitempty" json:"InheritTransformComponent,omitempty"` 
	VerletPhysicsComponent []VerletPhysicsComponent_Struct  `xml:"VerletPhysicsComponent,omitempty" json:"VerletPhysicsComponent,omitempty"` 
	AbilityComponent []AbilityComponent_Struct  `xml:"AbilityComponent,omitempty" json:"AbilityComponent,omitempty"` 
	ItemAlchemyComponent []ItemAlchemyComponent_Struct  `xml:"ItemAlchemyComponent,omitempty" json:"ItemAlchemyComponent,omitempty"` 
	ItemComponent []ItemComponent_Struct  `xml:"ItemComponent,omitempty" json:"ItemComponent,omitempty"` 
	ManaReloaderComponent []ManaReloaderComponent_Struct  `xml:"ManaReloaderComponent,omitempty" json:"ManaReloaderComponent,omitempty"` 
	MaterialAreaCheckerComponent []MaterialAreaCheckerComponent_Struct  `xml:"MaterialAreaCheckerComponent,omitempty" json:"MaterialAreaCheckerComponent,omitempty"` 
	SimplePhysicsComponent []SimplePhysicsComponent_Struct  `xml:"SimplePhysicsComponent,omitempty" json:"SimplePhysicsComponent,omitempty"` 
	ItemActionComponent []ItemActionComponent_Struct  `xml:"ItemActionComponent,omitempty" json:"ItemActionComponent,omitempty"` 
	SpriteOffsetAnimatorComponent []SpriteOffsetAnimatorComponent_Struct  `xml:"SpriteOffsetAnimatorComponent,omitempty" json:"SpriteOffsetAnimatorComponent,omitempty"` 
	ElectricChargeComponent []ElectricChargeComponent_Struct  `xml:"ElectricChargeComponent,omitempty" json:"ElectricChargeComponent,omitempty"` 
	ItemAIKnowledgeComponent []ItemAIKnowledgeComponent_Struct  `xml:"ItemAIKnowledgeComponent,omitempty" json:"ItemAIKnowledgeComponent,omitempty"` 
	UIInfoComponent []UIInfoComponent_Struct  `xml:"UIInfoComponent,omitempty" json:"UIInfoComponent,omitempty"` 
	ExplodeOnDamageComponent []ExplodeOnDamageComponent_Struct  `xml:"ExplodeOnDamageComponent,omitempty" json:"ExplodeOnDamageComponent,omitempty"` 
	PhysicsBodyComponent []PhysicsBodyComponent_Struct  `xml:"PhysicsBodyComponent,omitempty" json:"PhysicsBodyComponent,omitempty"` 
	PhysicsImageShapeComponent []PhysicsImageShapeComponent_Struct  `xml:"PhysicsImageShapeComponent,omitempty" json:"PhysicsImageShapeComponent,omitempty"` 
	PhysicsThrowableComponent []PhysicsThrowableComponent_Struct  `xml:"PhysicsThrowableComponent,omitempty" json:"PhysicsThrowableComponent,omitempty"` 
	PotionComponent []PotionComponent_Struct  `xml:"PotionComponent,omitempty" json:"PotionComponent,omitempty"` 
	PhysicsBodyCollisionDamageComponent []PhysicsBodyCollisionDamageComponent_Struct  `xml:"PhysicsBodyCollisionDamageComponent,omitempty" json:"PhysicsBodyCollisionDamageComponent,omitempty"` 
	GameEffectComponent []GameEffectComponent_Struct  `xml:"GameEffectComponent,omitempty" json:"GameEffectComponent,omitempty"` 
	UIIconComponent []UIIconComponent_Struct  `xml:"UIIconComponent,omitempty" json:"UIIconComponent,omitempty"` 
	MagicConvertMaterialComponent []MagicConvertMaterialComponent_Struct  `xml:"MagicConvertMaterialComponent,omitempty" json:"MagicConvertMaterialComponent,omitempty"` 
	EnergyShieldComponent []EnergyShieldComponent_Struct  `xml:"EnergyShieldComponent,omitempty" json:"EnergyShieldComponent,omitempty"` 
	LaserEmitterComponent []LaserEmitterComponent_Struct  `xml:"LaserEmitterComponent,omitempty" json:"LaserEmitterComponent,omitempty"` 
	CellEaterComponent []CellEaterComponent_Struct  `xml:"CellEaterComponent,omitempty" json:"CellEaterComponent,omitempty"` 
	WorldStateComponent []WorldStateComponent_Struct  `xml:"WorldStateComponent,omitempty" json:"WorldStateComponent,omitempty"` 
	PlayerStatsComponent []PlayerStatsComponent_Struct  `xml:"PlayerStatsComponent,omitempty" json:"PlayerStatsComponent,omitempty"` 
	TeleportComponent []TeleportComponent_Struct  `xml:"TeleportComponent,omitempty" json:"TeleportComponent,omitempty"` 
	LifetimeComponent []LifetimeComponent_Struct  `xml:"LifetimeComponent,omitempty" json:"LifetimeComponent,omitempty"` 
}

type Base_Struct struct {
	File string  `xml:"file,attr,omitempty" json:"file,omitempty"` 
	VelocityComponent []VelocityComponent_Struct  `xml:"VelocityComponent,omitempty" json:"VelocityComponent,omitempty"` 
}

type VelocityComponent_Struct struct {
	Gravity_y string  `xml:"gravity_y,attr,omitempty" json:"gravity_y,omitempty"` 
	Mass string  `xml:"mass,attr,omitempty" json:"mass,omitempty"` 
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Affect_physics_bodies string  `xml:"affect_physics_bodies,attr,omitempty" json:"affect_physics_bodies,omitempty"` 
	Air_friction string  `xml:"air_friction,attr,omitempty" json:"air_friction,omitempty"` 
	Apply_terminal_velocity string  `xml:"apply_terminal_velocity,attr,omitempty" json:"apply_terminal_velocity,omitempty"` 
	Displace_liquid string  `xml:"displace_liquid,attr,omitempty" json:"displace_liquid,omitempty"` 
	Gravity_x string  `xml:"gravity_x,attr,omitempty" json:"gravity_x,omitempty"` 
	Limit_to_max_velocity string  `xml:"limit_to_max_velocity,attr,omitempty" json:"limit_to_max_velocity,omitempty"` 
	Liquid_death_threshold string  `xml:"liquid_death_threshold,attr,omitempty" json:"liquid_death_threshold,omitempty"` 
	Liquid_drag string  `xml:"liquid_drag,attr,omitempty" json:"liquid_drag,omitempty"` 
	MVelocity_x string  `xml:"mVelocity.x,attr,omitempty" json:"mVelocity.x,omitempty"` 
	MVelocity_y string  `xml:"mVelocity.y,attr,omitempty" json:"mVelocity.y,omitempty"` 
	Terminal_velocity string  `xml:"terminal_velocity,attr,omitempty" json:"terminal_velocity,omitempty"` 
	Updates_velocity string  `xml:"updates_velocity,attr,omitempty" json:"updates_velocity,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
}

type ProjectileComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Speed_min string  `xml:"speed_min,attr,omitempty" json:"speed_min,omitempty"` 
	Speed_max string  `xml:"speed_max,attr,omitempty" json:"speed_max,omitempty"` 
	On_death_explode string  `xml:"on_death_explode,attr,omitempty" json:"on_death_explode,omitempty"` 
	On_death_gfx_leave_sprite string  `xml:"on_death_gfx_leave_sprite,attr,omitempty" json:"on_death_gfx_leave_sprite,omitempty"` 
	On_lifetime_out_explode string  `xml:"on_lifetime_out_explode,attr,omitempty" json:"on_lifetime_out_explode,omitempty"` 
	On_collision_die string  `xml:"on_collision_die,attr,omitempty" json:"on_collision_die,omitempty"` 
	Die_on_low_velocity string  `xml:"die_on_low_velocity,attr,omitempty" json:"die_on_low_velocity,omitempty"` 
	Damage string  `xml:"damage,attr,omitempty" json:"damage,omitempty"` 
	Lifetime string  `xml:"lifetime,attr,omitempty" json:"lifetime,omitempty"` 
	Shoot_light_flash_r string  `xml:"shoot_light_flash_r,attr,omitempty" json:"shoot_light_flash_r,omitempty"` 
	Shoot_light_flash_g string  `xml:"shoot_light_flash_g,attr,omitempty" json:"shoot_light_flash_g,omitempty"` 
	Shoot_light_flash_b string  `xml:"shoot_light_flash_b,attr,omitempty" json:"shoot_light_flash_b,omitempty"` 
	Shoot_light_flash_radius string  `xml:"shoot_light_flash_radius,attr,omitempty" json:"shoot_light_flash_radius,omitempty"` 
	Knockback_force string  `xml:"knockback_force,attr,omitempty" json:"knockback_force,omitempty"` 
	Physics_impulse_coeff string  `xml:"physics_impulse_coeff,attr,omitempty" json:"physics_impulse_coeff,omitempty"` 
	Config_explosion []Config_explosion_Struct  `xml:"config_explosion,omitempty" json:"config_explosion,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Angular_velocity string  `xml:"angular_velocity,attr,omitempty" json:"angular_velocity,omitempty"` 
	Attach_to_parent_trigger string  `xml:"attach_to_parent_trigger,attr,omitempty" json:"attach_to_parent_trigger,omitempty"` 
	Blood_count_multiplier string  `xml:"blood_count_multiplier,attr,omitempty" json:"blood_count_multiplier,omitempty"` 
	Bounce_always string  `xml:"bounce_always,attr,omitempty" json:"bounce_always,omitempty"` 
	Bounce_at_any_angle string  `xml:"bounce_at_any_angle,attr,omitempty" json:"bounce_at_any_angle,omitempty"` 
	Bounce_energy string  `xml:"bounce_energy,attr,omitempty" json:"bounce_energy,omitempty"` 
	Bounces_left string  `xml:"bounces_left,attr,omitempty" json:"bounces_left,omitempty"` 
	Camera_shake_when_shot string  `xml:"camera_shake_when_shot,attr,omitempty" json:"camera_shake_when_shot,omitempty"` 
	Collect_materials_to_shooter string  `xml:"collect_materials_to_shooter,attr,omitempty" json:"collect_materials_to_shooter,omitempty"` 
	Collide_with_entities string  `xml:"collide_with_entities,attr,omitempty" json:"collide_with_entities,omitempty"` 
	Collide_with_shooter_frames string  `xml:"collide_with_shooter_frames,attr,omitempty" json:"collide_with_shooter_frames,omitempty"` 
	Collide_with_tag string  `xml:"collide_with_tag,attr,omitempty" json:"collide_with_tag,omitempty"` 
	Collide_with_world string  `xml:"collide_with_world,attr,omitempty" json:"collide_with_world,omitempty"` 
	Create_shell_casing string  `xml:"create_shell_casing,attr,omitempty" json:"create_shell_casing,omitempty"` 
	Damage_every_x_frames string  `xml:"damage_every_x_frames,attr,omitempty" json:"damage_every_x_frames,omitempty"` 
	Damage_scale_max_speed string  `xml:"damage_scale_max_speed,attr,omitempty" json:"damage_scale_max_speed,omitempty"` 
	Damage_scaled_by_speed string  `xml:"damage_scaled_by_speed,attr,omitempty" json:"damage_scaled_by_speed,omitempty"` 
	Die_on_liquid_collision string  `xml:"die_on_liquid_collision,attr,omitempty" json:"die_on_liquid_collision,omitempty"` 
	Die_on_low_velocity_limit string  `xml:"die_on_low_velocity_limit,attr,omitempty" json:"die_on_low_velocity_limit,omitempty"` 
	Direction_nonrandom_rad string  `xml:"direction_nonrandom_rad,attr,omitempty" json:"direction_nonrandom_rad,omitempty"` 
	Direction_random_rad string  `xml:"direction_random_rad,attr,omitempty" json:"direction_random_rad,omitempty"` 
	Do_moveto_update string  `xml:"do_moveto_update,attr,omitempty" json:"do_moveto_update,omitempty"` 
	Explosion_dont_damage_shooter string  `xml:"explosion_dont_damage_shooter,attr,omitempty" json:"explosion_dont_damage_shooter,omitempty"` 
	Friction string  `xml:"friction,attr,omitempty" json:"friction,omitempty"` 
	Friendly_fire string  `xml:"friendly_fire,attr,omitempty" json:"friendly_fire,omitempty"` 
	Ground_collision_fx string  `xml:"ground_collision_fx,attr,omitempty" json:"ground_collision_fx,omitempty"` 
	Ground_penetration_coeff string  `xml:"ground_penetration_coeff,attr,omitempty" json:"ground_penetration_coeff,omitempty"` 
	Ground_penetration_max_durability_to_destroy string  `xml:"ground_penetration_max_durability_to_destroy,attr,omitempty" json:"ground_penetration_max_durability_to_destroy,omitempty"` 
	Hit_particle_force_multiplier string  `xml:"hit_particle_force_multiplier,attr,omitempty" json:"hit_particle_force_multiplier,omitempty"` 
	Lifetime_randomness string  `xml:"lifetime_randomness,attr,omitempty" json:"lifetime_randomness,omitempty"` 
	Lob_max string  `xml:"lob_max,attr,omitempty" json:"lob_max,omitempty"` 
	Lob_min string  `xml:"lob_min,attr,omitempty" json:"lob_min,omitempty"` 
	MLastFrameDamaged string  `xml:"mLastFrameDamaged,attr,omitempty" json:"mLastFrameDamaged,omitempty"` 
	Never_hit_player string  `xml:"never_hit_player,attr,omitempty" json:"never_hit_player,omitempty"` 
	On_collision_remove_projectile string  `xml:"on_collision_remove_projectile,attr,omitempty" json:"on_collision_remove_projectile,omitempty"` 
	On_collision_spawn_entity string  `xml:"on_collision_spawn_entity,attr,omitempty" json:"on_collision_spawn_entity,omitempty"` 
	On_death_duplicate_remaining string  `xml:"on_death_duplicate_remaining,attr,omitempty" json:"on_death_duplicate_remaining,omitempty"` 
	On_death_emit_particle string  `xml:"on_death_emit_particle,attr,omitempty" json:"on_death_emit_particle,omitempty"` 
	On_death_emit_particle_count string  `xml:"on_death_emit_particle_count,attr,omitempty" json:"on_death_emit_particle_count,omitempty"` 
	On_death_item_pickable_radius string  `xml:"on_death_item_pickable_radius,attr,omitempty" json:"on_death_item_pickable_radius,omitempty"` 
	On_death_particle_check_concrete string  `xml:"on_death_particle_check_concrete,attr,omitempty" json:"on_death_particle_check_concrete,omitempty"` 
	Penetrate_entities string  `xml:"penetrate_entities,attr,omitempty" json:"penetrate_entities,omitempty"` 
	Penetrate_world string  `xml:"penetrate_world,attr,omitempty" json:"penetrate_world,omitempty"` 
	Penetrate_world_velocity_coeff string  `xml:"penetrate_world_velocity_coeff,attr,omitempty" json:"penetrate_world_velocity_coeff,omitempty"` 
	Play_damage_sounds string  `xml:"play_damage_sounds,attr,omitempty" json:"play_damage_sounds,omitempty"` 
	Projectile_type string  `xml:"projectile_type,attr,omitempty" json:"projectile_type,omitempty"` 
	Ragdoll_force_multiplier string  `xml:"ragdoll_force_multiplier,attr,omitempty" json:"ragdoll_force_multiplier,omitempty"` 
	Ragdoll_fx_on_collision string  `xml:"ragdoll_fx_on_collision,attr,omitempty" json:"ragdoll_fx_on_collision,omitempty"` 
	Shell_casing_material string  `xml:"shell_casing_material,attr,omitempty" json:"shell_casing_material,omitempty"` 
	Shell_casing_offset_x string  `xml:"shell_casing_offset.x,attr,omitempty" json:"shell_casing_offset.x,omitempty"` 
	Shell_casing_offset_y string  `xml:"shell_casing_offset.y,attr,omitempty" json:"shell_casing_offset.y,omitempty"` 
	Spawn_entity_is_projectile string  `xml:"spawn_entity_is_projectile,attr,omitempty" json:"spawn_entity_is_projectile,omitempty"` 
	Velocity_sets_rotation string  `xml:"velocity_sets_rotation,attr,omitempty" json:"velocity_sets_rotation,omitempty"` 
	Velocity_sets_scale string  `xml:"velocity_sets_scale,attr,omitempty" json:"velocity_sets_scale,omitempty"` 
	Velocity_sets_scale_coeff string  `xml:"velocity_sets_scale_coeff,attr,omitempty" json:"velocity_sets_scale_coeff,omitempty"` 
	Velocity_sets_y_flip string  `xml:"velocity_sets_y_flip,attr,omitempty" json:"velocity_sets_y_flip,omitempty"` 
	Velocity_updates_animation string  `xml:"velocity_updates_animation,attr,omitempty" json:"velocity_updates_animation,omitempty"` 
	Config []Config_Struct  `xml:"config,omitempty" json:"config,omitempty"` 
	Damage_by_type []Damage_by_type_Struct  `xml:"damage_by_type,omitempty" json:"damage_by_type,omitempty"` 
	Damage_critical []Damage_critical_Struct  `xml:"damage_critical,omitempty" json:"damage_critical,omitempty"` 
	Muzzle_flash_file string  `xml:"muzzle_flash_file,attr,omitempty" json:"muzzle_flash_file,omitempty"` 
}

type Config_explosion_Struct struct {
	Never_cache string  `xml:"never_cache,attr,omitempty" json:"never_cache,omitempty"` 
	Damage string  `xml:"damage,attr,omitempty" json:"damage,omitempty"` 
	Camera_shake string  `xml:"camera_shake,attr,omitempty" json:"camera_shake,omitempty"` 
	Explosion_radius string  `xml:"explosion_radius,attr,omitempty" json:"explosion_radius,omitempty"` 
	Explosion_sprite string  `xml:"explosion_sprite,attr,omitempty" json:"explosion_sprite,omitempty"` 
	Explosion_sprite_lifetime string  `xml:"explosion_sprite_lifetime,attr,omitempty" json:"explosion_sprite_lifetime,omitempty"` 
	Create_cell_material string  `xml:"create_cell_material,attr,omitempty" json:"create_cell_material,omitempty"` 
	Create_cell_probability string  `xml:"create_cell_probability,attr,omitempty" json:"create_cell_probability,omitempty"` 
	Hole_destroy_liquid string  `xml:"hole_destroy_liquid,attr,omitempty" json:"hole_destroy_liquid,omitempty"` 
	Hole_enabled string  `xml:"hole_enabled,attr,omitempty" json:"hole_enabled,omitempty"` 
	Hole_image string  `xml:"hole_image,attr,omitempty" json:"hole_image,omitempty"` 
	Particle_effect string  `xml:"particle_effect,attr,omitempty" json:"particle_effect,omitempty"` 
	Damage_mortals string  `xml:"damage_mortals,attr,omitempty" json:"damage_mortals,omitempty"` 
	Physics_explosion_power_min string  `xml:"physics_explosion_power.min,attr,omitempty" json:"physics_explosion_power.min,omitempty"` 
	Physics_explosion_power_max string  `xml:"physics_explosion_power.max,attr,omitempty" json:"physics_explosion_power.max,omitempty"` 
	Physics_throw_enabled string  `xml:"physics_throw_enabled,attr,omitempty" json:"physics_throw_enabled,omitempty"` 
	Shake_vegetation string  `xml:"shake_vegetation,attr,omitempty" json:"shake_vegetation,omitempty"` 
	Sparks_count_max string  `xml:"sparks_count_max,attr,omitempty" json:"sparks_count_max,omitempty"` 
	Sparks_count_min string  `xml:"sparks_count_min,attr,omitempty" json:"sparks_count_min,omitempty"` 
	Sparks_enabled string  `xml:"sparks_enabled,attr,omitempty" json:"sparks_enabled,omitempty"` 
	Stains_enabled string  `xml:"stains_enabled,attr,omitempty" json:"stains_enabled,omitempty"` 
	Stains_radius string  `xml:"stains_radius,attr,omitempty" json:"stains_radius,omitempty"` 
	Audio_enabled string  `xml:"audio_enabled,attr,omitempty" json:"audio_enabled,omitempty"` 
	Audio_liquid_amount_normalized string  `xml:"audio_liquid_amount_normalized,attr,omitempty" json:"audio_liquid_amount_normalized,omitempty"` 
	Background_lightning_count string  `xml:"background_lightning_count,attr,omitempty" json:"background_lightning_count,omitempty"` 
	Cell_explosion_damage_required string  `xml:"cell_explosion_damage_required,attr,omitempty" json:"cell_explosion_damage_required,omitempty"` 
	Cell_explosion_power string  `xml:"cell_explosion_power,attr,omitempty" json:"cell_explosion_power,omitempty"` 
	Cell_explosion_power_ragdoll_coeff string  `xml:"cell_explosion_power_ragdoll_coeff,attr,omitempty" json:"cell_explosion_power_ragdoll_coeff,omitempty"` 
	Cell_explosion_probability string  `xml:"cell_explosion_probability,attr,omitempty" json:"cell_explosion_probability,omitempty"` 
	Cell_explosion_radius_max string  `xml:"cell_explosion_radius_max,attr,omitempty" json:"cell_explosion_radius_max,omitempty"` 
	Cell_explosion_radius_min string  `xml:"cell_explosion_radius_min,attr,omitempty" json:"cell_explosion_radius_min,omitempty"` 
	Cell_explosion_velocity_min string  `xml:"cell_explosion_velocity_min,attr,omitempty" json:"cell_explosion_velocity_min,omitempty"` 
	Crack_count string  `xml:"crack_count,attr,omitempty" json:"crack_count,omitempty"` 
	Delay_max string  `xml:"delay.max,attr,omitempty" json:"delay.max,omitempty"` 
	Delay_min string  `xml:"delay.min,attr,omitempty" json:"delay.min,omitempty"` 
	Destroy_non_platform_solid_enabled string  `xml:"destroy_non_platform_solid_enabled,attr,omitempty" json:"destroy_non_platform_solid_enabled,omitempty"` 
	Electricity_count string  `xml:"electricity_count,attr,omitempty" json:"electricity_count,omitempty"` 
	Explosion_delay_id string  `xml:"explosion_delay_id,attr,omitempty" json:"explosion_delay_id,omitempty"` 
	Explosion_sprite_additive string  `xml:"explosion_sprite_additive,attr,omitempty" json:"explosion_sprite_additive,omitempty"` 
	Explosion_sprite_emissive string  `xml:"explosion_sprite_emissive,attr,omitempty" json:"explosion_sprite_emissive,omitempty"` 
	Explosion_sprite_random_rotation string  `xml:"explosion_sprite_random_rotation,attr,omitempty" json:"explosion_sprite_random_rotation,omitempty"` 
	Gore_particle_count string  `xml:"gore_particle_count,attr,omitempty" json:"gore_particle_count,omitempty"` 
	Hole_destroy_physics_dynamic string  `xml:"hole_destroy_physics_dynamic,attr,omitempty" json:"hole_destroy_physics_dynamic,omitempty"` 
	Is_digger string  `xml:"is_digger,attr,omitempty" json:"is_digger,omitempty"` 
	Knockback_force string  `xml:"knockback_force,attr,omitempty" json:"knockback_force,omitempty"` 
	Light_b string  `xml:"light_b,attr,omitempty" json:"light_b,omitempty"` 
	Light_enabled string  `xml:"light_enabled,attr,omitempty" json:"light_enabled,omitempty"` 
	Light_fade_time string  `xml:"light_fade_time,attr,omitempty" json:"light_fade_time,omitempty"` 
	Light_g string  `xml:"light_g,attr,omitempty" json:"light_g,omitempty"` 
	Light_r string  `xml:"light_r,attr,omitempty" json:"light_r,omitempty"` 
	Light_radius_coeff string  `xml:"light_radius_coeff,attr,omitempty" json:"light_radius_coeff,omitempty"` 
	Material_sparks_count_max string  `xml:"material_sparks_count_max,attr,omitempty" json:"material_sparks_count_max,omitempty"` 
	Material_sparks_count_min string  `xml:"material_sparks_count_min,attr,omitempty" json:"material_sparks_count_min,omitempty"` 
	Material_sparks_enabled string  `xml:"material_sparks_enabled,attr,omitempty" json:"material_sparks_enabled,omitempty"` 
	Material_sparks_min_hp string  `xml:"material_sparks_min_hp,attr,omitempty" json:"material_sparks_min_hp,omitempty"` 
	Material_sparks_probability string  `xml:"material_sparks_probability,attr,omitempty" json:"material_sparks_probability,omitempty"` 
	Material_sparks_real string  `xml:"material_sparks_real,attr,omitempty" json:"material_sparks_real,omitempty"` 
	Material_sparks_scale_with_hp string  `xml:"material_sparks_scale_with_hp,attr,omitempty" json:"material_sparks_scale_with_hp,omitempty"` 
	Max_durability_to_destroy string  `xml:"max_durability_to_destroy,attr,omitempty" json:"max_durability_to_destroy,omitempty"` 
	Min_radius_for_cracks string  `xml:"min_radius_for_cracks,attr,omitempty" json:"min_radius_for_cracks,omitempty"` 
	Physics_multiplier_ragdoll_force string  `xml:"physics_multiplier_ragdoll_force,attr,omitempty" json:"physics_multiplier_ragdoll_force,omitempty"` 
	Pixel_sprites_enabled string  `xml:"pixel_sprites_enabled,attr,omitempty" json:"pixel_sprites_enabled,omitempty"` 
	Ray_energy string  `xml:"ray_energy,attr,omitempty" json:"ray_energy,omitempty"` 
	Spark_material string  `xml:"spark_material,attr,omitempty" json:"spark_material,omitempty"` 
	Sparks_inner_radius_coeff string  `xml:"sparks_inner_radius_coeff,attr,omitempty" json:"sparks_inner_radius_coeff,omitempty"` 
	Damage_critical []Damage_critical_Struct  `xml:"damage_critical,omitempty" json:"damage_critical,omitempty"` 
}

type SpriteComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Alpha string  `xml:"alpha,attr,omitempty" json:"alpha,omitempty"` 
	Image_file string  `xml:"image_file,attr,omitempty" json:"image_file,omitempty"` 
	Offset_x string  `xml:"offset_x,attr,omitempty" json:"offset_x,omitempty"` 
	Offset_y string  `xml:"offset_y,attr,omitempty" json:"offset_y,omitempty"` 
	Update_transform_rotation string  `xml:"update_transform_rotation,attr,omitempty" json:"update_transform_rotation,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Additive string  `xml:"additive,attr,omitempty" json:"additive,omitempty"` 
	Emissive string  `xml:"emissive,attr,omitempty" json:"emissive,omitempty"` 
	Fog_of_war_hole string  `xml:"fog_of_war_hole,attr,omitempty" json:"fog_of_war_hole,omitempty"` 
	Has_special_scale string  `xml:"has_special_scale,attr,omitempty" json:"has_special_scale,omitempty"` 
	Is_text_sprite string  `xml:"is_text_sprite,attr,omitempty" json:"is_text_sprite,omitempty"` 
	Kill_entity_after_finished string  `xml:"kill_entity_after_finished,attr,omitempty" json:"kill_entity_after_finished,omitempty"` 
	Never_ragdollify_on_death string  `xml:"never_ragdollify_on_death,attr,omitempty" json:"never_ragdollify_on_death,omitempty"` 
	Offset_animator_offset_x string  `xml:"offset_animator_offset.x,attr,omitempty" json:"offset_animator_offset.x,omitempty"` 
	Offset_animator_offset_y string  `xml:"offset_animator_offset.y,attr,omitempty" json:"offset_animator_offset.y,omitempty"` 
	Rect_animation string  `xml:"rect_animation,attr,omitempty" json:"rect_animation,omitempty"` 
	Smooth_filtering string  `xml:"smooth_filtering,attr,omitempty" json:"smooth_filtering,omitempty"` 
	Special_scale_x string  `xml:"special_scale_x,attr,omitempty" json:"special_scale_x,omitempty"` 
	Special_scale_y string  `xml:"special_scale_y,attr,omitempty" json:"special_scale_y,omitempty"` 
	Transform_offset_x string  `xml:"transform_offset.x,attr,omitempty" json:"transform_offset.x,omitempty"` 
	Transform_offset_y string  `xml:"transform_offset.y,attr,omitempty" json:"transform_offset.y,omitempty"` 
	Ui_is_parent string  `xml:"ui_is_parent,attr,omitempty" json:"ui_is_parent,omitempty"` 
	Update_transform string  `xml:"update_transform,attr,omitempty" json:"update_transform,omitempty"` 
	Visible string  `xml:"visible,attr,omitempty" json:"visible,omitempty"` 
	Z_index string  `xml:"z_index,attr,omitempty" json:"z_index,omitempty"` 
	Next_rect_animation string  `xml:"next_rect_animation,attr,omitempty" json:"next_rect_animation,omitempty"` 
}

type ParticleEmitterComponent_Struct struct {
	Emitted_material_name string  `xml:"emitted_material_name,attr,omitempty" json:"emitted_material_name,omitempty"` 
	Count_min string  `xml:"count_min,attr,omitempty" json:"count_min,omitempty"` 
	Count_max string  `xml:"count_max,attr,omitempty" json:"count_max,omitempty"` 
	Offset_x string  `xml:"offset.x,attr,omitempty" json:"offset.x,omitempty"` 
	Offset_y string  `xml:"offset.y,attr,omitempty" json:"offset.y,omitempty"` 
	X_pos_offset_min string  `xml:"x_pos_offset_min,attr,omitempty" json:"x_pos_offset_min,omitempty"` 
	Y_pos_offset_min string  `xml:"y_pos_offset_min,attr,omitempty" json:"y_pos_offset_min,omitempty"` 
	X_pos_offset_max string  `xml:"x_pos_offset_max,attr,omitempty" json:"x_pos_offset_max,omitempty"` 
	Y_pos_offset_max string  `xml:"y_pos_offset_max,attr,omitempty" json:"y_pos_offset_max,omitempty"` 
	X_vel_min string  `xml:"x_vel_min,attr,omitempty" json:"x_vel_min,omitempty"` 
	X_vel_max string  `xml:"x_vel_max,attr,omitempty" json:"x_vel_max,omitempty"` 
	Y_vel_min string  `xml:"y_vel_min,attr,omitempty" json:"y_vel_min,omitempty"` 
	Y_vel_max string  `xml:"y_vel_max,attr,omitempty" json:"y_vel_max,omitempty"` 
	Is_trail string  `xml:"is_trail,attr,omitempty" json:"is_trail,omitempty"` 
	Trail_gap string  `xml:"trail_gap,attr,omitempty" json:"trail_gap,omitempty"` 
	Lifetime_min string  `xml:"lifetime_min,attr,omitempty" json:"lifetime_min,omitempty"` 
	Lifetime_max string  `xml:"lifetime_max,attr,omitempty" json:"lifetime_max,omitempty"` 
	Create_real_particles string  `xml:"create_real_particles,attr,omitempty" json:"create_real_particles,omitempty"` 
	Emit_cosmetic_particles string  `xml:"emit_cosmetic_particles,attr,omitempty" json:"emit_cosmetic_particles,omitempty"` 
	Emission_interval_min_frames string  `xml:"emission_interval_min_frames,attr,omitempty" json:"emission_interval_min_frames,omitempty"` 
	Emission_interval_max_frames string  `xml:"emission_interval_max_frames,attr,omitempty" json:"emission_interval_max_frames,omitempty"` 
	Is_emitting string  `xml:"is_emitting,attr,omitempty" json:"is_emitting,omitempty"` 
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Airflow_force string  `xml:"airflow_force,attr,omitempty" json:"airflow_force,omitempty"` 
	Airflow_scale string  `xml:"airflow_scale,attr,omitempty" json:"airflow_scale,omitempty"` 
	Airflow_time string  `xml:"airflow_time,attr,omitempty" json:"airflow_time,omitempty"` 
	Area_circle_radius_max string  `xml:"area_circle_radius.max,attr,omitempty" json:"area_circle_radius.max,omitempty"` 
	Area_circle_radius_min string  `xml:"area_circle_radius.min,attr,omitempty" json:"area_circle_radius.min,omitempty"` 
	Area_circle_sector_degrees string  `xml:"area_circle_sector_degrees,attr,omitempty" json:"area_circle_sector_degrees,omitempty"` 
	Attractor_force string  `xml:"attractor_force,attr,omitempty" json:"attractor_force,omitempty"` 
	B2_force string  `xml:"b2_force,attr,omitempty" json:"b2_force,omitempty"` 
	Collide_with_gas_and_fire string  `xml:"collide_with_gas_and_fire,attr,omitempty" json:"collide_with_gas_and_fire,omitempty"` 
	Collide_with_grid string  `xml:"collide_with_grid,attr,omitempty" json:"collide_with_grid,omitempty"` 
	Color string  `xml:"color,attr,omitempty" json:"color,omitempty"` 
	Color_is_based_on_pos string  `xml:"color_is_based_on_pos,attr,omitempty" json:"color_is_based_on_pos,omitempty"` 
	Cosmetic_force_create string  `xml:"cosmetic_force_create,attr,omitempty" json:"cosmetic_force_create,omitempty"` 
	Custom_style string  `xml:"custom_style,attr,omitempty" json:"custom_style,omitempty"` 
	Delay_frames string  `xml:"delay_frames,attr,omitempty" json:"delay_frames,omitempty"` 
	Direction_random_deg string  `xml:"direction_random_deg,attr,omitempty" json:"direction_random_deg,omitempty"` 
	Draw_as_long string  `xml:"draw_as_long,attr,omitempty" json:"draw_as_long,omitempty"` 
	Emission_chance string  `xml:"emission_chance,attr,omitempty" json:"emission_chance,omitempty"` 
	Emit_real_particles string  `xml:"emit_real_particles,attr,omitempty" json:"emit_real_particles,omitempty"` 
	Emitter_lifetime_frames string  `xml:"emitter_lifetime_frames,attr,omitempty" json:"emitter_lifetime_frames,omitempty"` 
	Fade_based_on_lifetime string  `xml:"fade_based_on_lifetime,attr,omitempty" json:"fade_based_on_lifetime,omitempty"` 
	Fire_cells_dont_ignite_damagemodel string  `xml:"fire_cells_dont_ignite_damagemodel,attr,omitempty" json:"fire_cells_dont_ignite_damagemodel,omitempty"` 
	Friction string  `xml:"friction,attr,omitempty" json:"friction,omitempty"` 
	Gravity_x string  `xml:"gravity.x,attr,omitempty" json:"gravity.x,omitempty"` 
	Gravity_y string  `xml:"gravity.y,attr,omitempty" json:"gravity.y,omitempty"` 
	Image_animation_emission_probability string  `xml:"image_animation_emission_probability,attr,omitempty" json:"image_animation_emission_probability,omitempty"` 
	Image_animation_loop string  `xml:"image_animation_loop,attr,omitempty" json:"image_animation_loop,omitempty"` 
	Image_animation_phase string  `xml:"image_animation_phase,attr,omitempty" json:"image_animation_phase,omitempty"` 
	Image_animation_raytrace_from_center string  `xml:"image_animation_raytrace_from_center,attr,omitempty" json:"image_animation_raytrace_from_center,omitempty"` 
	Image_animation_speed string  `xml:"image_animation_speed,attr,omitempty" json:"image_animation_speed,omitempty"` 
	Image_animation_use_entity_rotation string  `xml:"image_animation_use_entity_rotation,attr,omitempty" json:"image_animation_use_entity_rotation,omitempty"` 
	Particle_single_width string  `xml:"particle_single_width,attr,omitempty" json:"particle_single_width,omitempty"` 
	Render_back string  `xml:"render_back,attr,omitempty" json:"render_back,omitempty"` 
	Render_on_grid string  `xml:"render_on_grid,attr,omitempty" json:"render_on_grid,omitempty"` 
	Render_ultrabright string  `xml:"render_ultrabright,attr,omitempty" json:"render_ultrabright,omitempty"` 
	Set_magic_creation string  `xml:"set_magic_creation,attr,omitempty" json:"set_magic_creation,omitempty"` 
	Use_material_inventory string  `xml:"use_material_inventory,attr,omitempty" json:"use_material_inventory,omitempty"` 
	Velocity_always_away_from_center string  `xml:"velocity_always_away_from_center,attr,omitempty" json:"velocity_always_away_from_center,omitempty"` 
}

type HitboxComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Aabb_min_x string  `xml:"aabb_min_x,attr,omitempty" json:"aabb_min_x,omitempty"` 
	Aabb_max_x string  `xml:"aabb_max_x,attr,omitempty" json:"aabb_max_x,omitempty"` 
	Aabb_min_y string  `xml:"aabb_min_y,attr,omitempty" json:"aabb_min_y,omitempty"` 
	Aabb_max_y string  `xml:"aabb_max_y,attr,omitempty" json:"aabb_max_y,omitempty"` 
	Damage_multiplier string  `xml:"damage_multiplier,attr,omitempty" json:"damage_multiplier,omitempty"` 
	Is_enemy string  `xml:"is_enemy,attr,omitempty" json:"is_enemy,omitempty"` 
	Is_item string  `xml:"is_item,attr,omitempty" json:"is_item,omitempty"` 
	Is_player string  `xml:"is_player,attr,omitempty" json:"is_player,omitempty"` 
	Offset_x string  `xml:"offset.x,attr,omitempty" json:"offset.x,omitempty"` 
	Offset_y string  `xml:"offset.y,attr,omitempty" json:"offset.y,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
}

type DamageModelComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Hp string  `xml:"hp,attr,omitempty" json:"hp,omitempty"` 
	Fire_probability_of_ignition string  `xml:"fire_probability_of_ignition,attr,omitempty" json:"fire_probability_of_ignition,omitempty"` 
	Falling_damages string  `xml:"falling_damages,attr,omitempty" json:"falling_damages,omitempty"` 
	Air_needed string  `xml:"air_needed,attr,omitempty" json:"air_needed,omitempty"` 
	Materials_damage string  `xml:"materials_damage,attr,omitempty" json:"materials_damage,omitempty"` 
	Blood_material string  `xml:"blood_material,attr,omitempty" json:"blood_material,omitempty"` 
	Blood_multiplier string  `xml:"blood_multiplier,attr,omitempty" json:"blood_multiplier,omitempty"` 
	Create_ragdoll string  `xml:"create_ragdoll,attr,omitempty" json:"create_ragdoll,omitempty"` 
	Air_in_lungs string  `xml:"air_in_lungs,attr,omitempty" json:"air_in_lungs,omitempty"` 
	Air_in_lungs_max string  `xml:"air_in_lungs_max,attr,omitempty" json:"air_in_lungs_max,omitempty"` 
	Air_lack_of_damage string  `xml:"air_lack_of_damage,attr,omitempty" json:"air_lack_of_damage,omitempty"` 
	Blood_spray_create_some_cosmetic string  `xml:"blood_spray_create_some_cosmetic,attr,omitempty" json:"blood_spray_create_some_cosmetic,omitempty"` 
	Blood_spray_material string  `xml:"blood_spray_material,attr,omitempty" json:"blood_spray_material,omitempty"` 
	Blood_sprite_directional string  `xml:"blood_sprite_directional,attr,omitempty" json:"blood_sprite_directional,omitempty"` 
	Blood_sprite_large string  `xml:"blood_sprite_large,attr,omitempty" json:"blood_sprite_large,omitempty"` 
	Critical_damage_resistance string  `xml:"critical_damage_resistance,attr,omitempty" json:"critical_damage_resistance,omitempty"` 
	Drop_items_on_death string  `xml:"drop_items_on_death,attr,omitempty" json:"drop_items_on_death,omitempty"` 
	Falling_damage_damage_max string  `xml:"falling_damage_damage_max,attr,omitempty" json:"falling_damage_damage_max,omitempty"` 
	Falling_damage_damage_min string  `xml:"falling_damage_damage_min,attr,omitempty" json:"falling_damage_damage_min,omitempty"` 
	Falling_damage_height_max string  `xml:"falling_damage_height_max,attr,omitempty" json:"falling_damage_height_max,omitempty"` 
	Falling_damage_height_min string  `xml:"falling_damage_height_min,attr,omitempty" json:"falling_damage_height_min,omitempty"` 
	Fire_damage_amount string  `xml:"fire_damage_amount,attr,omitempty" json:"fire_damage_amount,omitempty"` 
	Fire_damage_ignited_amount string  `xml:"fire_damage_ignited_amount,attr,omitempty" json:"fire_damage_ignited_amount,omitempty"` 
	Fire_how_much_fire_generates string  `xml:"fire_how_much_fire_generates,attr,omitempty" json:"fire_how_much_fire_generates,omitempty"` 
	In_liquid_shooting_electrify_prob string  `xml:"in_liquid_shooting_electrify_prob,attr,omitempty" json:"in_liquid_shooting_electrify_prob,omitempty"` 
	Invincibility_frames string  `xml:"invincibility_frames,attr,omitempty" json:"invincibility_frames,omitempty"` 
	Is_on_fire string  `xml:"is_on_fire,attr,omitempty" json:"is_on_fire,omitempty"` 
	Kill_now string  `xml:"kill_now,attr,omitempty" json:"kill_now,omitempty"` 
	MLastElectricityResistanceFrame string  `xml:"mLastElectricityResistanceFrame,attr,omitempty" json:"mLastElectricityResistanceFrame,omitempty"` 
	MLastFrameReportedBlock string  `xml:"mLastFrameReportedBlock,attr,omitempty" json:"mLastFrameReportedBlock,omitempty"` 
	MLastMaxHpChangeFrame string  `xml:"mLastMaxHpChangeFrame,attr,omitempty" json:"mLastMaxHpChangeFrame,omitempty"` 
	Material_damage_min_cell_count string  `xml:"material_damage_min_cell_count,attr,omitempty" json:"material_damage_min_cell_count,omitempty"` 
	Materials_create_messages string  `xml:"materials_create_messages,attr,omitempty" json:"materials_create_messages,omitempty"` 
	Materials_damage_proportional_to_maxhp string  `xml:"materials_damage_proportional_to_maxhp,attr,omitempty" json:"materials_damage_proportional_to_maxhp,omitempty"` 
	Materials_how_much_damage string  `xml:"materials_how_much_damage,attr,omitempty" json:"materials_how_much_damage,omitempty"` 
	Materials_that_create_messages string  `xml:"materials_that_create_messages,attr,omitempty" json:"materials_that_create_messages,omitempty"` 
	Materials_that_damage string  `xml:"materials_that_damage,attr,omitempty" json:"materials_that_damage,omitempty"` 
	Max_hp string  `xml:"max_hp,attr,omitempty" json:"max_hp,omitempty"` 
	Max_hp_cap string  `xml:"max_hp_cap,attr,omitempty" json:"max_hp_cap,omitempty"` 
	Max_hp_old string  `xml:"max_hp_old,attr,omitempty" json:"max_hp_old,omitempty"` 
	Minimum_knockback_force string  `xml:"minimum_knockback_force,attr,omitempty" json:"minimum_knockback_force,omitempty"` 
	Physics_objects_damage string  `xml:"physics_objects_damage,attr,omitempty" json:"physics_objects_damage,omitempty"` 
	Ragdoll_blood_amount_absolute string  `xml:"ragdoll_blood_amount_absolute,attr,omitempty" json:"ragdoll_blood_amount_absolute,omitempty"` 
	Ragdoll_filenames_file string  `xml:"ragdoll_filenames_file,attr,omitempty" json:"ragdoll_filenames_file,omitempty"` 
	Ragdoll_fx_forced string  `xml:"ragdoll_fx_forced,attr,omitempty" json:"ragdoll_fx_forced,omitempty"` 
	Ragdoll_material string  `xml:"ragdoll_material,attr,omitempty" json:"ragdoll_material,omitempty"` 
	Ragdoll_offset_x string  `xml:"ragdoll_offset_x,attr,omitempty" json:"ragdoll_offset_x,omitempty"` 
	Ragdoll_offset_y string  `xml:"ragdoll_offset_y,attr,omitempty" json:"ragdoll_offset_y,omitempty"` 
	Ragdollify_child_entity_sprites string  `xml:"ragdollify_child_entity_sprites,attr,omitempty" json:"ragdollify_child_entity_sprites,omitempty"` 
	Ragdollify_disintegrate_nonroot string  `xml:"ragdollify_disintegrate_nonroot,attr,omitempty" json:"ragdollify_disintegrate_nonroot,omitempty"` 
	Ragdollify_root_angular_damping string  `xml:"ragdollify_root_angular_damping,attr,omitempty" json:"ragdollify_root_angular_damping,omitempty"` 
	Ui_force_report_damage string  `xml:"ui_force_report_damage,attr,omitempty" json:"ui_force_report_damage,omitempty"` 
	Ui_report_damage string  `xml:"ui_report_damage,attr,omitempty" json:"ui_report_damage,omitempty"` 
	Wait_for_kill_flag_on_death string  `xml:"wait_for_kill_flag_on_death,attr,omitempty" json:"wait_for_kill_flag_on_death,omitempty"` 
	Wet_status_effect_damage string  `xml:"wet_status_effect_damage,attr,omitempty" json:"wet_status_effect_damage,omitempty"` 
	Damage_multipliers []Damage_multipliers_Struct  `xml:"damage_multipliers,omitempty" json:"damage_multipliers,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
}

type AudioComponent_Struct struct {
	File string  `xml:"file,attr,omitempty" json:"file,omitempty"` 
	Event_root string  `xml:"event_root,attr,omitempty" json:"event_root,omitempty"` 
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Audio_physics_material string  `xml:"audio_physics_material,attr,omitempty" json:"audio_physics_material,omitempty"` 
	Play_only_if_visible string  `xml:"play_only_if_visible,attr,omitempty" json:"play_only_if_visible,omitempty"` 
	Remove_latest_event_on_destroyed string  `xml:"remove_latest_event_on_destroyed,attr,omitempty" json:"remove_latest_event_on_destroyed,omitempty"` 
	Send_message_on_event_dead string  `xml:"send_message_on_event_dead,attr,omitempty" json:"send_message_on_event_dead,omitempty"` 
	Set_latest_event_position string  `xml:"set_latest_event_position,attr,omitempty" json:"set_latest_event_position,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
}

type VariableStorageComponent_Struct struct {
	Name string  `xml:"name,attr,omitempty" json:"name,omitempty"` 
	Value_string string  `xml:"value_string,attr,omitempty" json:"value_string,omitempty"` 
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Value_bool string  `xml:"value_bool,attr,omitempty" json:"value_bool,omitempty"` 
	Value_float string  `xml:"value_float,attr,omitempty" json:"value_float,omitempty"` 
	Value_int string  `xml:"value_int,attr,omitempty" json:"value_int,omitempty"` 
}

type Transform_Struct struct {
	Position_x string  `xml:"position.x,attr,omitempty" json:"position.x,omitempty"` 
	Position_y string  `xml:"position.y,attr,omitempty" json:"position.y,omitempty"` 
	Rotation string  `xml:"rotation,attr,omitempty" json:"rotation,omitempty"` 
	Scale_x string  `xml:"scale.x,attr,omitempty" json:"scale.x,omitempty"` 
	Scale_y string  `xml:"scale.y,attr,omitempty" json:"scale.y,omitempty"` 
}

type AudioListenerComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Z string  `xml:"z,attr,omitempty" json:"z,omitempty"` 
}

type AudioLoopComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Auto_play string  `xml:"auto_play,attr,omitempty" json:"auto_play,omitempty"` 
	Auto_play_if_enabled string  `xml:"auto_play_if_enabled,attr,omitempty" json:"auto_play_if_enabled,omitempty"` 
	Calculate_material_lowpass string  `xml:"calculate_material_lowpass,attr,omitempty" json:"calculate_material_lowpass,omitempty"` 
	Event_name string  `xml:"event_name,attr,omitempty" json:"event_name,omitempty"` 
	File string  `xml:"file,attr,omitempty" json:"file,omitempty"` 
	Play_on_component_enable string  `xml:"play_on_component_enable,attr,omitempty" json:"play_on_component_enable,omitempty"` 
	Set_speed_parameter string  `xml:"set_speed_parameter,attr,omitempty" json:"set_speed_parameter,omitempty"` 
	Set_speed_parameter_only_based_on_x_movement string  `xml:"set_speed_parameter_only_based_on_x_movement,attr,omitempty" json:"set_speed_parameter_only_based_on_x_movement,omitempty"` 
	Set_speed_parameter_only_based_on_y_movement string  `xml:"set_speed_parameter_only_based_on_y_movement,attr,omitempty" json:"set_speed_parameter_only_based_on_y_movement,omitempty"` 
	Volume_autofade_speed string  `xml:"volume_autofade_speed,attr,omitempty" json:"volume_autofade_speed,omitempty"` 
}

type CharacterDataComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Buoyancy_check_offset_y string  `xml:"buoyancy_check_offset_y,attr,omitempty" json:"buoyancy_check_offset_y,omitempty"` 
	Check_collision_max_size_x string  `xml:"check_collision_max_size_x,attr,omitempty" json:"check_collision_max_size_x,omitempty"` 
	Check_collision_max_size_y string  `xml:"check_collision_max_size_y,attr,omitempty" json:"check_collision_max_size_y,omitempty"` 
	Climb_over_y string  `xml:"climb_over_y,attr,omitempty" json:"climb_over_y,omitempty"` 
	Collision_aabb_max_x string  `xml:"collision_aabb_max_x,attr,omitempty" json:"collision_aabb_max_x,omitempty"` 
	Collision_aabb_max_y string  `xml:"collision_aabb_max_y,attr,omitempty" json:"collision_aabb_max_y,omitempty"` 
	Collision_aabb_min_x string  `xml:"collision_aabb_min_x,attr,omitempty" json:"collision_aabb_min_x,omitempty"` 
	Collision_aabb_min_y string  `xml:"collision_aabb_min_y,attr,omitempty" json:"collision_aabb_min_y,omitempty"` 
	Destroy_ground string  `xml:"destroy_ground,attr,omitempty" json:"destroy_ground,omitempty"` 
	Eff_hg_b2force_multiplier string  `xml:"eff_hg_b2force_multiplier,attr,omitempty" json:"eff_hg_b2force_multiplier,omitempty"` 
	Eff_hg_damage_max string  `xml:"eff_hg_damage_max,attr,omitempty" json:"eff_hg_damage_max,omitempty"` 
	Eff_hg_damage_min string  `xml:"eff_hg_damage_min,attr,omitempty" json:"eff_hg_damage_min,omitempty"` 
	Eff_hg_offset_y string  `xml:"eff_hg_offset_y,attr,omitempty" json:"eff_hg_offset_y,omitempty"` 
	Eff_hg_position_x string  `xml:"eff_hg_position_x,attr,omitempty" json:"eff_hg_position_x,omitempty"` 
	Eff_hg_position_y string  `xml:"eff_hg_position_y,attr,omitempty" json:"eff_hg_position_y,omitempty"` 
	Eff_hg_size_x string  `xml:"eff_hg_size_x,attr,omitempty" json:"eff_hg_size_x,omitempty"` 
	Eff_hg_size_y string  `xml:"eff_hg_size_y,attr,omitempty" json:"eff_hg_size_y,omitempty"` 
	Eff_hg_update_box2d string  `xml:"eff_hg_update_box2d,attr,omitempty" json:"eff_hg_update_box2d,omitempty"` 
	Eff_hg_velocity_max_x string  `xml:"eff_hg_velocity_max_x,attr,omitempty" json:"eff_hg_velocity_max_x,omitempty"` 
	Eff_hg_velocity_max_y string  `xml:"eff_hg_velocity_max_y,attr,omitempty" json:"eff_hg_velocity_max_y,omitempty"` 
	Eff_hg_velocity_min_x string  `xml:"eff_hg_velocity_min_x,attr,omitempty" json:"eff_hg_velocity_min_x,omitempty"` 
	Eff_hg_velocity_min_y string  `xml:"eff_hg_velocity_min_y,attr,omitempty" json:"eff_hg_velocity_min_y,omitempty"` 
	Effect_hit_ground string  `xml:"effect_hit_ground,attr,omitempty" json:"effect_hit_ground,omitempty"` 
	Fly_recharge_spd string  `xml:"fly_recharge_spd,attr,omitempty" json:"fly_recharge_spd,omitempty"` 
	Fly_recharge_spd_ground string  `xml:"fly_recharge_spd_ground,attr,omitempty" json:"fly_recharge_spd_ground,omitempty"` 
	Fly_time_max string  `xml:"fly_time_max,attr,omitempty" json:"fly_time_max,omitempty"` 
	Flying_in_air_wait_frames string  `xml:"flying_in_air_wait_frames,attr,omitempty" json:"flying_in_air_wait_frames,omitempty"` 
	Flying_needs_recharge string  `xml:"flying_needs_recharge,attr,omitempty" json:"flying_needs_recharge,omitempty"` 
	Flying_recharge_removal_frames string  `xml:"flying_recharge_removal_frames,attr,omitempty" json:"flying_recharge_removal_frames,omitempty"` 
	Gravity string  `xml:"gravity,attr,omitempty" json:"gravity,omitempty"` 
	Ground_stickyness string  `xml:"ground_stickyness,attr,omitempty" json:"ground_stickyness,omitempty"` 
	Is_on_ground string  `xml:"is_on_ground,attr,omitempty" json:"is_on_ground,omitempty"` 
	Is_on_slippery_ground string  `xml:"is_on_slippery_ground,attr,omitempty" json:"is_on_slippery_ground,omitempty"` 
	Liquid_velocity_coeff string  `xml:"liquid_velocity_coeff,attr,omitempty" json:"liquid_velocity_coeff,omitempty"` 
	MFlyingTimeLeft string  `xml:"mFlyingTimeLeft,attr,omitempty" json:"mFlyingTimeLeft,omitempty"` 
	Mass string  `xml:"mass,attr,omitempty" json:"mass,omitempty"` 
	Platforming_type string  `xml:"platforming_type,attr,omitempty" json:"platforming_type,omitempty"` 
	Send_transform_update_message string  `xml:"send_transform_update_message,attr,omitempty" json:"send_transform_update_message,omitempty"` 
}

type CharacterPlatformingComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Accel_x string  `xml:"accel_x,attr,omitempty" json:"accel_x,omitempty"` 
	Accel_x_air string  `xml:"accel_x_air,attr,omitempty" json:"accel_x_air,omitempty"` 
	Fly_model_player string  `xml:"fly_model_player,attr,omitempty" json:"fly_model_player,omitempty"` 
	Fly_smooth_y string  `xml:"fly_smooth_y,attr,omitempty" json:"fly_smooth_y,omitempty"` 
	Fly_speed_change_spd string  `xml:"fly_speed_change_spd,attr,omitempty" json:"fly_speed_change_spd,omitempty"` 
	Fly_speed_max_down string  `xml:"fly_speed_max_down,attr,omitempty" json:"fly_speed_max_down,omitempty"` 
	Fly_speed_max_up string  `xml:"fly_speed_max_up,attr,omitempty" json:"fly_speed_max_up,omitempty"` 
	Fly_speed_mult string  `xml:"fly_speed_mult,attr,omitempty" json:"fly_speed_mult,omitempty"` 
	Fly_velocity_x string  `xml:"fly_velocity_x,attr,omitempty" json:"fly_velocity_x,omitempty"` 
	Jump_keydown_buffer string  `xml:"jump_keydown_buffer,attr,omitempty" json:"jump_keydown_buffer,omitempty"` 
	Jump_velocity_x string  `xml:"jump_velocity_x,attr,omitempty" json:"jump_velocity_x,omitempty"` 
	Jump_velocity_y string  `xml:"jump_velocity_y,attr,omitempty" json:"jump_velocity_y,omitempty"` 
	Keyboard_look string  `xml:"keyboard_look,attr,omitempty" json:"keyboard_look,omitempty"` 
	Mouse_look string  `xml:"mouse_look,attr,omitempty" json:"mouse_look,omitempty"` 
	Mouse_look_buffer string  `xml:"mouse_look_buffer,attr,omitempty" json:"mouse_look_buffer,omitempty"` 
	Pixel_gravity string  `xml:"pixel_gravity,attr,omitempty" json:"pixel_gravity,omitempty"` 
	Precision_jumping_max_duration_frames string  `xml:"precision_jumping_max_duration_frames,attr,omitempty" json:"precision_jumping_max_duration_frames,omitempty"` 
	Run_animation_velocity_switching_enabled string  `xml:"run_animation_velocity_switching_enabled,attr,omitempty" json:"run_animation_velocity_switching_enabled,omitempty"` 
	Run_animation_velocity_switching_threshold string  `xml:"run_animation_velocity_switching_threshold,attr,omitempty" json:"run_animation_velocity_switching_threshold,omitempty"` 
	Run_velocity string  `xml:"run_velocity,attr,omitempty" json:"run_velocity,omitempty"` 
	Swim_down_buoyancy_coeff string  `xml:"swim_down_buoyancy_coeff,attr,omitempty" json:"swim_down_buoyancy_coeff,omitempty"` 
	Swim_drag string  `xml:"swim_drag,attr,omitempty" json:"swim_drag,omitempty"` 
	Swim_extra_horizontal_drag string  `xml:"swim_extra_horizontal_drag,attr,omitempty" json:"swim_extra_horizontal_drag,omitempty"` 
	Swim_idle_buoyancy_coeff string  `xml:"swim_idle_buoyancy_coeff,attr,omitempty" json:"swim_idle_buoyancy_coeff,omitempty"` 
	Swim_up_buoyancy_coeff string  `xml:"swim_up_buoyancy_coeff,attr,omitempty" json:"swim_up_buoyancy_coeff,omitempty"` 
	Turn_animation_frames_between string  `xml:"turn_animation_frames_between,attr,omitempty" json:"turn_animation_frames_between,omitempty"` 
	Turning_buffer string  `xml:"turning_buffer,attr,omitempty" json:"turning_buffer,omitempty"` 
	Velocity_max_x string  `xml:"velocity_max_x,attr,omitempty" json:"velocity_max_x,omitempty"` 
	Velocity_max_y string  `xml:"velocity_max_y,attr,omitempty" json:"velocity_max_y,omitempty"` 
	Velocity_min_x string  `xml:"velocity_min_x,attr,omitempty" json:"velocity_min_x,omitempty"` 
	Velocity_min_y string  `xml:"velocity_min_y,attr,omitempty" json:"velocity_min_y,omitempty"` 
}

type ControlsComponent_Struct struct {
	//Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Enabled string  `xml:"enabled,attr,omitempty" json:"enabled,omitempty"` 
	Gamepad_fire_on_thumbstick_extend string  `xml:"gamepad_fire_on_thumbstick_extend,attr,omitempty" json:"gamepad_fire_on_thumbstick_extend,omitempty"` 
	Gamepad_fire_on_thumbstick_extend_threshold string  `xml:"gamepad_fire_on_thumbstick_extend_threshold,attr,omitempty" json:"gamepad_fire_on_thumbstick_extend_threshold,omitempty"` 
	Gamepad_indirect_aiming_enabled string  `xml:"gamepad_indirect_aiming_enabled,attr,omitempty" json:"gamepad_indirect_aiming_enabled,omitempty"` 
	Polymorph_hax string  `xml:"polymorph_hax,attr,omitempty" json:"polymorph_hax,omitempty"` 
	Polymorph_next_attack_frame string  `xml:"polymorph_next_attack_frame,attr,omitempty" json:"polymorph_next_attack_frame,omitempty"` 
}

type Damage_multipliers_Struct struct {
	Curse string  `xml:"curse,attr,omitempty" json:"curse,omitempty"` 
	Drill string  `xml:"drill,attr,omitempty" json:"drill,omitempty"` 
	Electricity string  `xml:"electricity,attr,omitempty" json:"electricity,omitempty"` 
	Explosion string  `xml:"explosion,attr,omitempty" json:"explosion,omitempty"` 
	Fire string  `xml:"fire,attr,omitempty" json:"fire,omitempty"` 
	Healing string  `xml:"healing,attr,omitempty" json:"healing,omitempty"` 
	Ice string  `xml:"ice,attr,omitempty" json:"ice,omitempty"` 
	Melee string  `xml:"melee,attr,omitempty" json:"melee,omitempty"` 
	Overeating string  `xml:"overeating,attr,omitempty" json:"overeating,omitempty"` 
	Physics_hit string  `xml:"physics_hit,attr,omitempty" json:"physics_hit,omitempty"` 
	Poison string  `xml:"poison,attr,omitempty" json:"poison,omitempty"` 
	Projectile string  `xml:"projectile,attr,omitempty" json:"projectile,omitempty"` 
	Radioactive string  `xml:"radioactive,attr,omitempty" json:"radioactive,omitempty"` 
	Slice string  `xml:"slice,attr,omitempty" json:"slice,omitempty"` 
}

type DrugEffectComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Drug_fx_target []Drug_fx_target_Struct  `xml:"drug_fx_target,omitempty" json:"drug_fx_target,omitempty"` 
	M_drug_fx_current []M_drug_fx_current_Struct  `xml:"m_drug_fx_current,omitempty" json:"m_drug_fx_current,omitempty"` 
}

type Drug_fx_target_Struct struct {
	Color_amount string  `xml:"color_amount,attr,omitempty" json:"color_amount,omitempty"` 
	Distortion_amount string  `xml:"distortion_amount,attr,omitempty" json:"distortion_amount,omitempty"` 
	Doublevision_amount string  `xml:"doublevision_amount,attr,omitempty" json:"doublevision_amount,omitempty"` 
	Fractals_amount string  `xml:"fractals_amount,attr,omitempty" json:"fractals_amount,omitempty"` 
	Fractals_size string  `xml:"fractals_size,attr,omitempty" json:"fractals_size,omitempty"` 
	Nightvision_amount string  `xml:"nightvision_amount,attr,omitempty" json:"nightvision_amount,omitempty"` 
}

type M_drug_fx_current_Struct struct {
	Color_amount string  `xml:"color_amount,attr,omitempty" json:"color_amount,omitempty"` 
	Distortion_amount string  `xml:"distortion_amount,attr,omitempty" json:"distortion_amount,omitempty"` 
	Doublevision_amount string  `xml:"doublevision_amount,attr,omitempty" json:"doublevision_amount,omitempty"` 
	Fractals_amount string  `xml:"fractals_amount,attr,omitempty" json:"fractals_amount,omitempty"` 
	Fractals_size string  `xml:"fractals_size,attr,omitempty" json:"fractals_size,omitempty"` 
	Nightvision_amount string  `xml:"nightvision_amount,attr,omitempty" json:"nightvision_amount,omitempty"` 
}

type GameLogComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Report_damage string  `xml:"report_damage,attr,omitempty" json:"report_damage,omitempty"` 
	Report_death string  `xml:"report_death,attr,omitempty" json:"report_death,omitempty"` 
	Report_new_biomes string  `xml:"report_new_biomes,attr,omitempty" json:"report_new_biomes,omitempty"` 
	MVisitiedBiomes []MVisitiedBiomes_Struct  `xml:"mVisitiedBiomes,omitempty" json:"mVisitiedBiomes,omitempty"` 
}

type MVisitiedBiomes_Struct struct {
	String_Struct []string  `xml:"string,omitempty" json:"string,omitempty"` 
}

type GameStatsComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Is_player string  `xml:"is_player,attr,omitempty" json:"is_player,omitempty"` 
	Name string  `xml:"name,attr,omitempty" json:"name,omitempty"` 
	Player_polymorph_count string  `xml:"player_polymorph_count,attr,omitempty" json:"player_polymorph_count,omitempty"` 
	Stats_filename string  `xml:"stats_filename,attr,omitempty" json:"stats_filename,omitempty"` 
}

type GenomeDataComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Berserk_dont_attack_friends string  `xml:"berserk_dont_attack_friends,attr,omitempty" json:"berserk_dont_attack_friends,omitempty"` 
	Food_chain_rank string  `xml:"food_chain_rank,attr,omitempty" json:"food_chain_rank,omitempty"` 
	Herd_id string  `xml:"herd_id,attr,omitempty" json:"herd_id,omitempty"` 
	Is_predator string  `xml:"is_predator,attr,omitempty" json:"is_predator,omitempty"` 
}

type GunComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
}

type HotspotComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Offset_x string  `xml:"offset.x,attr,omitempty" json:"offset.x,omitempty"` 
	Offset_y string  `xml:"offset.y,attr,omitempty" json:"offset.y,omitempty"` 
	Sprite_hotspot_name string  `xml:"sprite_hotspot_name,attr,omitempty" json:"sprite_hotspot_name,omitempty"` 
	Transform_with_scale string  `xml:"transform_with_scale,attr,omitempty" json:"transform_with_scale,omitempty"` 
}

type IngestionComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Blood_healing_speed string  `xml:"blood_healing_speed,attr,omitempty" json:"blood_healing_speed,omitempty"` 
	Ingestion_capacity string  `xml:"ingestion_capacity,attr,omitempty" json:"ingestion_capacity,omitempty"` 
	Ingestion_cooldown_delay_frames string  `xml:"ingestion_cooldown_delay_frames,attr,omitempty" json:"ingestion_cooldown_delay_frames,omitempty"` 
	Ingestion_reduce_every_n_frame string  `xml:"ingestion_reduce_every_n_frame,attr,omitempty" json:"ingestion_reduce_every_n_frame,omitempty"` 
	Ingestion_size string  `xml:"ingestion_size,attr,omitempty" json:"ingestion_size,omitempty"` 
	M_ingestion_cooldown_frames string  `xml:"m_ingestion_cooldown_frames,attr,omitempty" json:"m_ingestion_cooldown_frames,omitempty"` 
	Overingestion_damage string  `xml:"overingestion_damage,attr,omitempty" json:"overingestion_damage,omitempty"` 
}

type Inventory2Component_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Full_inventory_slots_x string  `xml:"full_inventory_slots_x,attr,omitempty" json:"full_inventory_slots_x,omitempty"` 
	Full_inventory_slots_y string  `xml:"full_inventory_slots_y,attr,omitempty" json:"full_inventory_slots_y,omitempty"` 
	MSavedActiveItemIndex string  `xml:"mSavedActiveItemIndex,attr,omitempty" json:"mSavedActiveItemIndex,omitempty"` 
	Quick_inventory_slots string  `xml:"quick_inventory_slots,attr,omitempty" json:"quick_inventory_slots,omitempty"` 
}

type InventoryGuiComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Has_opened_inventory_edit string  `xml:"has_opened_inventory_edit,attr,omitempty" json:"has_opened_inventory_edit,omitempty"` 
	Wallet_money_target string  `xml:"wallet_money_target,attr,omitempty" json:"wallet_money_target,omitempty"` 
}

type ItemPickUpperComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Drop_items_on_death string  `xml:"drop_items_on_death,attr,omitempty" json:"drop_items_on_death,omitempty"` 
	Is_immune_to_kicks string  `xml:"is_immune_to_kicks,attr,omitempty" json:"is_immune_to_kicks,omitempty"` 
	Is_in_npc string  `xml:"is_in_npc,attr,omitempty" json:"is_in_npc,omitempty"` 
	MLatestItemOverlapInfoBoxPosition_x string  `xml:"mLatestItemOverlapInfoBoxPosition.x,attr,omitempty" json:"mLatestItemOverlapInfoBoxPosition.x,omitempty"` 
	MLatestItemOverlapInfoBoxPosition_y string  `xml:"mLatestItemOverlapInfoBoxPosition.y,attr,omitempty" json:"mLatestItemOverlapInfoBoxPosition.y,omitempty"` 
	Only_pick_this_entity string  `xml:"only_pick_this_entity,attr,omitempty" json:"only_pick_this_entity,omitempty"` 
	Pick_up_any_item_buggy string  `xml:"pick_up_any_item_buggy,attr,omitempty" json:"pick_up_any_item_buggy,omitempty"` 
}

type KickComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Can_kick string  `xml:"can_kick,attr,omitempty" json:"can_kick,omitempty"` 
	Kick_damage string  `xml:"kick_damage,attr,omitempty" json:"kick_damage,omitempty"` 
	Kick_entities string  `xml:"kick_entities,attr,omitempty" json:"kick_entities,omitempty"` 
	Kick_knockback string  `xml:"kick_knockback,attr,omitempty" json:"kick_knockback,omitempty"` 
	Kick_radius string  `xml:"kick_radius,attr,omitempty" json:"kick_radius,omitempty"` 
	Max_force string  `xml:"max_force,attr,omitempty" json:"max_force,omitempty"` 
	Player_kickforce string  `xml:"player_kickforce,attr,omitempty" json:"player_kickforce,omitempty"` 
	Telekinesis_throw_speed string  `xml:"telekinesis_throw_speed,attr,omitempty" json:"telekinesis_throw_speed,omitempty"` 
}

type LightComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	B string  `xml:"b,attr,omitempty" json:"b,omitempty"` 
	Blinking_freq string  `xml:"blinking_freq,attr,omitempty" json:"blinking_freq,omitempty"` 
	Fade_out_time string  `xml:"fade_out_time,attr,omitempty" json:"fade_out_time,omitempty"` 
	G string  `xml:"g,attr,omitempty" json:"g,omitempty"` 
	Offset_x string  `xml:"offset_x,attr,omitempty" json:"offset_x,omitempty"` 
	Offset_y string  `xml:"offset_y,attr,omitempty" json:"offset_y,omitempty"` 
	R string  `xml:"r,attr,omitempty" json:"r,omitempty"` 
	Radius string  `xml:"radius,attr,omitempty" json:"radius,omitempty"` 
	Update_properties string  `xml:"update_properties,attr,omitempty" json:"update_properties,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
}

type LiquidDisplacerComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Radius string  `xml:"radius,attr,omitempty" json:"radius,omitempty"` 
	Velocity_x string  `xml:"velocity_x,attr,omitempty" json:"velocity_x,omitempty"` 
	Velocity_y string  `xml:"velocity_y,attr,omitempty" json:"velocity_y,omitempty"` 
}

type LuaComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Call_init_function string  `xml:"call_init_function,attr,omitempty" json:"call_init_function,omitempty"` 
	Enable_coroutines string  `xml:"enable_coroutines,attr,omitempty" json:"enable_coroutines,omitempty"` 
	Execute_every_n_frame string  `xml:"execute_every_n_frame,attr,omitempty" json:"execute_every_n_frame,omitempty"` 
	Execute_on_added string  `xml:"execute_on_added,attr,omitempty" json:"execute_on_added,omitempty"` 
	Execute_on_removed string  `xml:"execute_on_removed,attr,omitempty" json:"execute_on_removed,omitempty"` 
	Execute_times string  `xml:"execute_times,attr,omitempty" json:"execute_times,omitempty"` 
	MLastExecutionFrame string  `xml:"mLastExecutionFrame,attr,omitempty" json:"mLastExecutionFrame,omitempty"` 
	MModAppendsDone string  `xml:"mModAppendsDone,attr,omitempty" json:"mModAppendsDone,omitempty"` 
	Remove_after_executed string  `xml:"remove_after_executed,attr,omitempty" json:"remove_after_executed,omitempty"` 
	Script_source_file string  `xml:"script_source_file,attr,omitempty" json:"script_source_file,omitempty"` 
	Vm_type string  `xml:"vm_type,attr,omitempty" json:"vm_type,omitempty"` 
	Script_polymorphing_to string  `xml:"script_polymorphing_to,attr,omitempty" json:"script_polymorphing_to,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Script_wand_fired string  `xml:"script_wand_fired,attr,omitempty" json:"script_wand_fired,omitempty"` 
	Script_item_picked_up string  `xml:"script_item_picked_up,attr,omitempty" json:"script_item_picked_up,omitempty"` 
	Script_material_area_checker_success string  `xml:"script_material_area_checker_success,attr,omitempty" json:"script_material_area_checker_success,omitempty"` 
	Script_death string  `xml:"script_death,attr,omitempty" json:"script_death,omitempty"` 
}

type MaterialInventoryComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Audio_collision_size_modifier_amount string  `xml:"audio_collision_size_modifier_amount,attr,omitempty" json:"audio_collision_size_modifier_amount,omitempty"` 
	B2_force_on_leak string  `xml:"b2_force_on_leak,attr,omitempty" json:"b2_force_on_leak,omitempty"` 
	Death_throw_particle_velocity_coeff string  `xml:"death_throw_particle_velocity_coeff,attr,omitempty" json:"death_throw_particle_velocity_coeff,omitempty"` 
	Drop_as_item string  `xml:"drop_as_item,attr,omitempty" json:"drop_as_item,omitempty"` 
	Halftime_materials string  `xml:"halftime_materials,attr,omitempty" json:"halftime_materials,omitempty"` 
	Kill_when_empty string  `xml:"kill_when_empty,attr,omitempty" json:"kill_when_empty,omitempty"` 
	Last_frame_drank string  `xml:"last_frame_drank,attr,omitempty" json:"last_frame_drank,omitempty"` 
	Leak_on_damage_percent string  `xml:"leak_on_damage_percent,attr,omitempty" json:"leak_on_damage_percent,omitempty"` 
	Leak_pressure_max string  `xml:"leak_pressure_max,attr,omitempty" json:"leak_pressure_max,omitempty"` 
	Leak_pressure_min string  `xml:"leak_pressure_min,attr,omitempty" json:"leak_pressure_min,omitempty"` 
	Max_capacity string  `xml:"max_capacity,attr,omitempty" json:"max_capacity,omitempty"` 
	Min_damage_to_leak string  `xml:"min_damage_to_leak,attr,omitempty" json:"min_damage_to_leak,omitempty"` 
	On_death_spill string  `xml:"on_death_spill,attr,omitempty" json:"on_death_spill,omitempty"` 
	Count_per_material_type []Count_per_material_type_Struct  `xml:"count_per_material_type,omitempty" json:"count_per_material_type,omitempty"` 
}

type Count_per_material_type_Struct struct {
	Material []Material_Struct  `xml:"Material,omitempty" json:"Material,omitempty"` 
}

type Material_Struct struct {
	Count string  `xml:"count,attr,omitempty" json:"count,omitempty"` 
	Material string  `xml:"material,attr,omitempty" json:"material,omitempty"` 
}

type MaterialSuckerComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Barrel_size string  `xml:"barrel_size,attr,omitempty" json:"barrel_size,omitempty"` 
	Last_material_id string  `xml:"last_material_id,attr,omitempty" json:"last_material_id,omitempty"` 
	MAmountUsed string  `xml:"mAmountUsed,attr,omitempty" json:"mAmountUsed,omitempty"` 
	Material_type string  `xml:"material_type,attr,omitempty" json:"material_type,omitempty"` 
	Num_cells_sucked_per_frame string  `xml:"num_cells_sucked_per_frame,attr,omitempty" json:"num_cells_sucked_per_frame,omitempty"` 
	Randomized_position_max_x string  `xml:"randomized_position.max_x,attr,omitempty" json:"randomized_position.max_x,omitempty"` 
	Randomized_position_max_y string  `xml:"randomized_position.max_y,attr,omitempty" json:"randomized_position.max_y,omitempty"` 
	Randomized_position_min_x string  `xml:"randomized_position.min_x,attr,omitempty" json:"randomized_position.min_x,omitempty"` 
	Randomized_position_min_y string  `xml:"randomized_position.min_y,attr,omitempty" json:"randomized_position.min_y,omitempty"` 
	Set_projectile_to_liquid string  `xml:"set_projectile_to_liquid,attr,omitempty" json:"set_projectile_to_liquid,omitempty"` 
	Suck_gold string  `xml:"suck_gold,attr,omitempty" json:"suck_gold,omitempty"` 
	Suck_health string  `xml:"suck_health,attr,omitempty" json:"suck_health,omitempty"` 
	Suck_static_materials string  `xml:"suck_static_materials,attr,omitempty" json:"suck_static_materials,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
}

type PathFindingGridMarkerComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Marker_offset_x string  `xml:"marker_offset_x,attr,omitempty" json:"marker_offset_x,omitempty"` 
	Marker_offset_y string  `xml:"marker_offset_y,attr,omitempty" json:"marker_offset_y,omitempty"` 
	Marker_work_flag string  `xml:"marker_work_flag,attr,omitempty" json:"marker_work_flag,omitempty"` 
	Player_marker_radius string  `xml:"player_marker_radius,attr,omitempty" json:"player_marker_radius,omitempty"` 
}

type PhysicsPickUpComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Original_left_joint_pos_x string  `xml:"original_left_joint_pos.x,attr,omitempty" json:"original_left_joint_pos.x,omitempty"` 
	Original_left_joint_pos_y string  `xml:"original_left_joint_pos.y,attr,omitempty" json:"original_left_joint_pos.y,omitempty"` 
	Original_right_joint_pos_x string  `xml:"original_right_joint_pos.x,attr,omitempty" json:"original_right_joint_pos.x,omitempty"` 
	Original_right_joint_pos_y string  `xml:"original_right_joint_pos.y,attr,omitempty" json:"original_right_joint_pos.y,omitempty"` 
	Pick_up_strength string  `xml:"pick_up_strength,attr,omitempty" json:"pick_up_strength,omitempty"` 
	Transform []Transform_Struct  `xml:"transform,omitempty" json:"transform,omitempty"` 
}

type PlatformShooterPlayerComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Aiming_reticle_distance_from_character string  `xml:"aiming_reticle_distance_from_character,attr,omitempty" json:"aiming_reticle_distance_from_character,omitempty"` 
	Alcohol_drunken_speed string  `xml:"alcohol_drunken_speed,attr,omitempty" json:"alcohol_drunken_speed,omitempty"` 
	Blood_fungi_drunken_speed string  `xml:"blood_fungi_drunken_speed,attr,omitempty" json:"blood_fungi_drunken_speed,omitempty"` 
	Blood_worm_drunken_speed string  `xml:"blood_worm_drunken_speed,attr,omitempty" json:"blood_worm_drunken_speed,omitempty"` 
	Camera_max_distance_from_character string  `xml:"camera_max_distance_from_character,attr,omitempty" json:"camera_max_distance_from_character,omitempty"` 
	Center_camera_on_this_entity string  `xml:"center_camera_on_this_entity,attr,omitempty" json:"center_camera_on_this_entity,omitempty"` 
	Eating_area_max_x string  `xml:"eating_area_max.x,attr,omitempty" json:"eating_area_max.x,omitempty"` 
	Eating_area_max_y string  `xml:"eating_area_max.y,attr,omitempty" json:"eating_area_max.y,omitempty"` 
	Eating_area_min_x string  `xml:"eating_area_min.x,attr,omitempty" json:"eating_area_min.x,omitempty"` 
	Eating_area_min_y string  `xml:"eating_area_min.y,attr,omitempty" json:"eating_area_min.y,omitempty"` 
	Eating_cells_per_frame string  `xml:"eating_cells_per_frame,attr,omitempty" json:"eating_cells_per_frame,omitempty"` 
	Eating_delay_frames string  `xml:"eating_delay_frames,attr,omitempty" json:"eating_delay_frames,omitempty"` 
	Eating_probability string  `xml:"eating_probability,attr,omitempty" json:"eating_probability,omitempty"` 
	Move_camera_with_aim string  `xml:"move_camera_with_aim,attr,omitempty" json:"move_camera_with_aim,omitempty"` 
	Stoned_speed string  `xml:"stoned_speed,attr,omitempty" json:"stoned_speed,omitempty"` 
}

type PlayerCollisionComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Getting_crushed_threshold string  `xml:"getting_crushed_threshold,attr,omitempty" json:"getting_crushed_threshold,omitempty"` 
	Moving_up_before_getting_crushed_threshold string  `xml:"moving_up_before_getting_crushed_threshold,attr,omitempty" json:"moving_up_before_getting_crushed_threshold,omitempty"` 
}

type ShotEffectComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Condition_effect string  `xml:"condition_effect,attr,omitempty" json:"condition_effect,omitempty"` 
	Condition_status string  `xml:"condition_status,attr,omitempty" json:"condition_status,omitempty"` 
	Extra_modifier string  `xml:"extra_modifier,attr,omitempty" json:"extra_modifier,omitempty"` 
}

type SpriteAnimatorComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Rotate_to_surface_normal string  `xml:"rotate_to_surface_normal,attr,omitempty" json:"rotate_to_surface_normal,omitempty"` 
	Target_sprite_comp_name string  `xml:"target_sprite_comp_name,attr,omitempty" json:"target_sprite_comp_name,omitempty"` 
}

type SpriteParticleEmitterComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Additive string  `xml:"additive,attr,omitempty" json:"additive,omitempty"` 
	Angular_velocity string  `xml:"angular_velocity,attr,omitempty" json:"angular_velocity,omitempty"` 
	Camera_bound string  `xml:"camera_bound,attr,omitempty" json:"camera_bound,omitempty"` 
	Camera_distance string  `xml:"camera_distance,attr,omitempty" json:"camera_distance,omitempty"` 
	Color_a string  `xml:"color.a,attr,omitempty" json:"color.a,omitempty"` 
	Color_b string  `xml:"color.b,attr,omitempty" json:"color.b,omitempty"` 
	Color_g string  `xml:"color.g,attr,omitempty" json:"color.g,omitempty"` 
	Color_r string  `xml:"color.r,attr,omitempty" json:"color.r,omitempty"` 
	Color_change_a string  `xml:"color_change.a,attr,omitempty" json:"color_change.a,omitempty"` 
	Color_change_b string  `xml:"color_change.b,attr,omitempty" json:"color_change.b,omitempty"` 
	Color_change_g string  `xml:"color_change.g,attr,omitempty" json:"color_change.g,omitempty"` 
	Color_change_r string  `xml:"color_change.r,attr,omitempty" json:"color_change.r,omitempty"` 
	Count_max string  `xml:"count_max,attr,omitempty" json:"count_max,omitempty"` 
	Count_min string  `xml:"count_min,attr,omitempty" json:"count_min,omitempty"` 
	Delay string  `xml:"delay,attr,omitempty" json:"delay,omitempty"` 
	Emission_interval_max_frames string  `xml:"emission_interval_max_frames,attr,omitempty" json:"emission_interval_max_frames,omitempty"` 
	Emission_interval_min_frames string  `xml:"emission_interval_min_frames,attr,omitempty" json:"emission_interval_min_frames,omitempty"` 
	Emissive string  `xml:"emissive,attr,omitempty" json:"emissive,omitempty"` 
	Entity_velocity_multiplier string  `xml:"entity_velocity_multiplier,attr,omitempty" json:"entity_velocity_multiplier,omitempty"` 
	Expand_randomize_position_x string  `xml:"expand_randomize_position.x,attr,omitempty" json:"expand_randomize_position.x,omitempty"` 
	Expand_randomize_position_y string  `xml:"expand_randomize_position.y,attr,omitempty" json:"expand_randomize_position.y,omitempty"` 
	Gravity_x string  `xml:"gravity.x,attr,omitempty" json:"gravity.x,omitempty"` 
	Gravity_y string  `xml:"gravity.y,attr,omitempty" json:"gravity.y,omitempty"` 
	Is_emitting string  `xml:"is_emitting,attr,omitempty" json:"is_emitting,omitempty"` 
	Lifetime string  `xml:"lifetime,attr,omitempty" json:"lifetime,omitempty"` 
	Randomize_alpha_max string  `xml:"randomize_alpha.max,attr,omitempty" json:"randomize_alpha.max,omitempty"` 
	Randomize_alpha_min string  `xml:"randomize_alpha.min,attr,omitempty" json:"randomize_alpha.min,omitempty"` 
	Randomize_angular_velocity_max string  `xml:"randomize_angular_velocity.max,attr,omitempty" json:"randomize_angular_velocity.max,omitempty"` 
	Randomize_angular_velocity_min string  `xml:"randomize_angular_velocity.min,attr,omitempty" json:"randomize_angular_velocity.min,omitempty"` 
	Randomize_animation_speed_coeff_max string  `xml:"randomize_animation_speed_coeff.max,attr,omitempty" json:"randomize_animation_speed_coeff.max,omitempty"` 
	Randomize_animation_speed_coeff_min string  `xml:"randomize_animation_speed_coeff.min,attr,omitempty" json:"randomize_animation_speed_coeff.min,omitempty"` 
	Randomize_lifetime_max string  `xml:"randomize_lifetime.max,attr,omitempty" json:"randomize_lifetime.max,omitempty"` 
	Randomize_lifetime_min string  `xml:"randomize_lifetime.min,attr,omitempty" json:"randomize_lifetime.min,omitempty"` 
	Randomize_position_max_x string  `xml:"randomize_position.max_x,attr,omitempty" json:"randomize_position.max_x,omitempty"` 
	Randomize_position_max_y string  `xml:"randomize_position.max_y,attr,omitempty" json:"randomize_position.max_y,omitempty"` 
	Randomize_position_min_x string  `xml:"randomize_position.min_x,attr,omitempty" json:"randomize_position.min_x,omitempty"` 
	Randomize_position_min_y string  `xml:"randomize_position.min_y,attr,omitempty" json:"randomize_position.min_y,omitempty"` 
	Randomize_position_inside_hitbox string  `xml:"randomize_position_inside_hitbox,attr,omitempty" json:"randomize_position_inside_hitbox,omitempty"` 
	Randomize_rotation_max string  `xml:"randomize_rotation.max,attr,omitempty" json:"randomize_rotation.max,omitempty"` 
	Randomize_rotation_min string  `xml:"randomize_rotation.min,attr,omitempty" json:"randomize_rotation.min,omitempty"` 
	Randomize_scale_max_x string  `xml:"randomize_scale.max_x,attr,omitempty" json:"randomize_scale.max_x,omitempty"` 
	Randomize_scale_max_y string  `xml:"randomize_scale.max_y,attr,omitempty" json:"randomize_scale.max_y,omitempty"` 
	Randomize_scale_min_x string  `xml:"randomize_scale.min_x,attr,omitempty" json:"randomize_scale.min_x,omitempty"` 
	Randomize_scale_min_y string  `xml:"randomize_scale.min_y,attr,omitempty" json:"randomize_scale.min_y,omitempty"` 
	Randomize_velocity_max_x string  `xml:"randomize_velocity.max_x,attr,omitempty" json:"randomize_velocity.max_x,omitempty"` 
	Randomize_velocity_max_y string  `xml:"randomize_velocity.max_y,attr,omitempty" json:"randomize_velocity.max_y,omitempty"` 
	Randomize_velocity_min_x string  `xml:"randomize_velocity.min_x,attr,omitempty" json:"randomize_velocity.min_x,omitempty"` 
	Randomize_velocity_min_y string  `xml:"randomize_velocity.min_y,attr,omitempty" json:"randomize_velocity.min_y,omitempty"` 
	Render_back string  `xml:"render_back,attr,omitempty" json:"render_back,omitempty"` 
	Rotation string  `xml:"rotation,attr,omitempty" json:"rotation,omitempty"` 
	Scale_x string  `xml:"scale.x,attr,omitempty" json:"scale.x,omitempty"` 
	Scale_y string  `xml:"scale.y,attr,omitempty" json:"scale.y,omitempty"` 
	Scale_velocity_x string  `xml:"scale_velocity.x,attr,omitempty" json:"scale_velocity.x,omitempty"` 
	Scale_velocity_y string  `xml:"scale_velocity.y,attr,omitempty" json:"scale_velocity.y,omitempty"` 
	Sprite_centered string  `xml:"sprite_centered,attr,omitempty" json:"sprite_centered,omitempty"` 
	Sprite_file string  `xml:"sprite_file,attr,omitempty" json:"sprite_file,omitempty"` 
	Sprite_random_rotation string  `xml:"sprite_random_rotation,attr,omitempty" json:"sprite_random_rotation,omitempty"` 
	Use_rotation_from_entity string  `xml:"use_rotation_from_entity,attr,omitempty" json:"use_rotation_from_entity,omitempty"` 
	Use_rotation_from_velocity_component string  `xml:"use_rotation_from_velocity_component,attr,omitempty" json:"use_rotation_from_velocity_component,omitempty"` 
	Use_velocity_as_rotation string  `xml:"use_velocity_as_rotation,attr,omitempty" json:"use_velocity_as_rotation,omitempty"` 
	Velocity_x string  `xml:"velocity.x,attr,omitempty" json:"velocity.x,omitempty"` 
	Velocity_y string  `xml:"velocity.y,attr,omitempty" json:"velocity.y,omitempty"` 
	Velocity_always_away_from_center string  `xml:"velocity_always_away_from_center,attr,omitempty" json:"velocity_always_away_from_center,omitempty"` 
	Velocity_slowdown string  `xml:"velocity_slowdown,attr,omitempty" json:"velocity_slowdown,omitempty"` 
	Z_index string  `xml:"z_index,attr,omitempty" json:"z_index,omitempty"` 
}

type SpriteStainsComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Fade_stains_towards_srite_top string  `xml:"fade_stains_towards_srite_top,attr,omitempty" json:"fade_stains_towards_srite_top,omitempty"` 
	Sprite_id string  `xml:"sprite_id,attr,omitempty" json:"sprite_id,omitempty"` 
	Stain_shaken_drop_chance_multiplier string  `xml:"stain_shaken_drop_chance_multiplier,attr,omitempty" json:"stain_shaken_drop_chance_multiplier,omitempty"` 
	MData []MData_Struct  `xml:"mData,omitempty" json:"mData,omitempty"` 
}

type MData_Struct struct {
	Reference_frame_h string  `xml:"reference_frame_h,attr,omitempty" json:"reference_frame_h,omitempty"` 
	Reference_frame_materials string  `xml:"reference_frame_materials,attr,omitempty" json:"reference_frame_materials,omitempty"` 
	Reference_frame_stains string  `xml:"reference_frame_stains,attr,omitempty" json:"reference_frame_stains,omitempty"` 
	Reference_frame_w string  `xml:"reference_frame_w,attr,omitempty" json:"reference_frame_w,omitempty"` 
}

type StatusEffectDataComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Ingestion_effect_causes string  `xml:"ingestion_effect_causes,attr,omitempty" json:"ingestion_effect_causes,omitempty"` 
	Stain_effects []Stain_effects_Struct  `xml:"stain_effects,omitempty" json:"stain_effects,omitempty"` 
	Effects_previous []Effects_previous_Struct  `xml:"effects_previous,omitempty" json:"effects_previous,omitempty"` 
	Ingestion_effects []Ingestion_effects_Struct  `xml:"ingestion_effects,omitempty" json:"ingestion_effects,omitempty"` 
	Ingestion_effect_causes_many []Ingestion_effect_causes_many_Struct  `xml:"ingestion_effect_causes_many,omitempty" json:"ingestion_effect_causes_many,omitempty"` 
}

type Stain_effects_Struct struct {
	Primitive_Struct []string  `xml:"primitive,omitempty" json:"primitive,omitempty"` 
}

type Effects_previous_Struct struct {
	Primitive_Struct []string  `xml:"primitive,omitempty" json:"primitive,omitempty"` 
}

type Ingestion_effects_Struct struct {
	Primitive_Struct []string  `xml:"primitive,omitempty" json:"primitive,omitempty"` 
}

type Ingestion_effect_causes_many_Struct struct {
	Primitive_Struct []string  `xml:"primitive,omitempty" json:"primitive,omitempty"` 
}

type StreamingKeepAliveComponent_Struct struct {
	TEMP_TEMPY string  `xml:"TEMP_TEMPY,attr,omitempty" json:"TEMP_TEMPY,omitempty"` 
	TEMP_TEMP_TEMP string  `xml:"TEMP_TEMP_TEMP,attr,omitempty" json:"TEMP_TEMP_TEMP,omitempty"` 
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
}

type TelekinesisComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Kick_to_use string  `xml:"kick_to_use,attr,omitempty" json:"kick_to_use,omitempty"` 
	MBodyID string  `xml:"mBodyID,attr,omitempty" json:"mBodyID,omitempty"` 
	MInteract string  `xml:"mInteract,attr,omitempty" json:"mInteract,omitempty"` 
	MMinBodyDistance string  `xml:"mMinBodyDistance,attr,omitempty" json:"mMinBodyDistance,omitempty"` 
	MStartAimAngle string  `xml:"mStartAimAngle,attr,omitempty" json:"mStartAimAngle,omitempty"` 
	MStartBodyAngle string  `xml:"mStartBodyAngle,attr,omitempty" json:"mStartBodyAngle,omitempty"` 
	MStartBodyDistance string  `xml:"mStartBodyDistance,attr,omitempty" json:"mStartBodyDistance,omitempty"` 
	MStartBodyMaxExtent string  `xml:"mStartBodyMaxExtent,attr,omitempty" json:"mStartBodyMaxExtent,omitempty"` 
	MStartTime string  `xml:"mStartTime,attr,omitempty" json:"mStartTime,omitempty"` 
	MState string  `xml:"mState,attr,omitempty" json:"mState,omitempty"` 
	Max_size string  `xml:"max_size,attr,omitempty" json:"max_size,omitempty"` 
	Min_size string  `xml:"min_size,attr,omitempty" json:"min_size,omitempty"` 
	Radius string  `xml:"radius,attr,omitempty" json:"radius,omitempty"` 
	Target_distance string  `xml:"target_distance,attr,omitempty" json:"target_distance,omitempty"` 
	Throw_speed string  `xml:"throw_speed,attr,omitempty" json:"throw_speed,omitempty"` 
}

type WalletComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	MHasReachedInf string  `xml:"mHasReachedInf,attr,omitempty" json:"mHasReachedInf,omitempty"` 
	MMoneyPrevFrame string  `xml:"mMoneyPrevFrame,attr,omitempty" json:"mMoneyPrevFrame,omitempty"` 
	Money string  `xml:"money,attr,omitempty" json:"money,omitempty"` 
	Money_spent string  `xml:"money_spent,attr,omitempty" json:"money_spent,omitempty"` 
}

type InheritTransformComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Always_use_immediate_parent_rotation string  `xml:"always_use_immediate_parent_rotation,attr,omitempty" json:"always_use_immediate_parent_rotation,omitempty"` 
	Only_position string  `xml:"only_position,attr,omitempty" json:"only_position,omitempty"` 
	Parent_hotspot_tag string  `xml:"parent_hotspot_tag,attr,omitempty" json:"parent_hotspot_tag,omitempty"` 
	Parent_sprite_id string  `xml:"parent_sprite_id,attr,omitempty" json:"parent_sprite_id,omitempty"` 
	Rotate_based_on_x_scale string  `xml:"rotate_based_on_x_scale,attr,omitempty" json:"rotate_based_on_x_scale,omitempty"` 
	Use_root_parent string  `xml:"use_root_parent,attr,omitempty" json:"use_root_parent,omitempty"` 
	Transform []Transform_Struct  `xml:"Transform,omitempty" json:"Transform,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
}

type VerletPhysicsComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Animation_amount string  `xml:"animation_amount,attr,omitempty" json:"animation_amount,omitempty"` 
	Animation_energy string  `xml:"animation_energy,attr,omitempty" json:"animation_energy,omitempty"` 
	Animation_speed string  `xml:"animation_speed,attr,omitempty" json:"animation_speed,omitempty"` 
	Animation_target_offset_x string  `xml:"animation_target_offset.x,attr,omitempty" json:"animation_target_offset.x,omitempty"` 
	Animation_target_offset_y string  `xml:"animation_target_offset.y,attr,omitempty" json:"animation_target_offset.y,omitempty"` 
	Cloth_color string  `xml:"cloth_color,attr,omitempty" json:"cloth_color,omitempty"` 
	Cloth_color_edge string  `xml:"cloth_color_edge,attr,omitempty" json:"cloth_color_edge,omitempty"` 
	Cloth_sprite_z_index string  `xml:"cloth_sprite_z_index,attr,omitempty" json:"cloth_sprite_z_index,omitempty"` 
	Collide_with_cells string  `xml:"collide_with_cells,attr,omitempty" json:"collide_with_cells,omitempty"` 
	Colors string  `xml:"colors,attr,omitempty" json:"colors,omitempty"` 
	Constrain_stretching string  `xml:"constrain_stretching,attr,omitempty" json:"constrain_stretching,omitempty"` 
	Follow_entity_transform string  `xml:"follow_entity_transform,attr,omitempty" json:"follow_entity_transform,omitempty"` 
	Gets_entity_velocity_coeff string  `xml:"gets_entity_velocity_coeff,attr,omitempty" json:"gets_entity_velocity_coeff,omitempty"` 
	M_is_culled_previous string  `xml:"m_is_culled_previous,attr,omitempty" json:"m_is_culled_previous,omitempty"` 
	M_position_previous_x string  `xml:"m_position_previous.x,attr,omitempty" json:"m_position_previous.x,omitempty"` 
	M_position_previous_y string  `xml:"m_position_previous.y,attr,omitempty" json:"m_position_previous.y,omitempty"` 
	Mass_max string  `xml:"mass_max,attr,omitempty" json:"mass_max,omitempty"` 
	Mass_min string  `xml:"mass_min,attr,omitempty" json:"mass_min,omitempty"` 
	Materials string  `xml:"materials,attr,omitempty" json:"materials,omitempty"` 
	Num_links string  `xml:"num_links,attr,omitempty" json:"num_links,omitempty"` 
	Num_points string  `xml:"num_points,attr,omitempty" json:"num_points,omitempty"` 
	Pixelate_sprite_transforms string  `xml:"pixelate_sprite_transforms,attr,omitempty" json:"pixelate_sprite_transforms,omitempty"` 
	Resting_distance string  `xml:"resting_distance,attr,omitempty" json:"resting_distance,omitempty"` 
	Scale_sprite_x string  `xml:"scale_sprite_x,attr,omitempty" json:"scale_sprite_x,omitempty"` 
	Simulate_gravity string  `xml:"simulate_gravity,attr,omitempty" json:"simulate_gravity,omitempty"` 
	Simulate_wind string  `xml:"simulate_wind,attr,omitempty" json:"simulate_wind,omitempty"` 
	Stain_cells_probability string  `xml:"stain_cells_probability,attr,omitempty" json:"stain_cells_probability,omitempty"` 
	Stiffness string  `xml:"stiffness,attr,omitempty" json:"stiffness,omitempty"` 
	Type string  `xml:"type,attr,omitempty" json:"type,omitempty"` 
	Velocity_dampening string  `xml:"velocity_dampening,attr,omitempty" json:"velocity_dampening,omitempty"` 
	Width string  `xml:"width,attr,omitempty" json:"width,omitempty"` 
	Wind_change_speed string  `xml:"wind_change_speed,attr,omitempty" json:"wind_change_speed,omitempty"` 
}

type AbilityComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Amount_in_inventory string  `xml:"amount_in_inventory,attr,omitempty" json:"amount_in_inventory,omitempty"` 
	Base_item_file string  `xml:"base_item_file,attr,omitempty" json:"base_item_file,omitempty"` 
	Charge_wait_frames string  `xml:"charge_wait_frames,attr,omitempty" json:"charge_wait_frames,omitempty"` 
	Click_to_use string  `xml:"click_to_use,attr,omitempty" json:"click_to_use,omitempty"` 
	Cooldown_frames string  `xml:"cooldown_frames,attr,omitempty" json:"cooldown_frames,omitempty"` 
	Current_slot_durability string  `xml:"current_slot_durability,attr,omitempty" json:"current_slot_durability,omitempty"` 
	Drop_as_item_on_death string  `xml:"drop_as_item_on_death,attr,omitempty" json:"drop_as_item_on_death,omitempty"` 
	Entity_count string  `xml:"entity_count,attr,omitempty" json:"entity_count,omitempty"` 
	Fast_projectile string  `xml:"fast_projectile,attr,omitempty" json:"fast_projectile,omitempty"` 
	Gun_level string  `xml:"gun_level,attr,omitempty" json:"gun_level,omitempty"` 
	Is_petris_gun string  `xml:"is_petris_gun,attr,omitempty" json:"is_petris_gun,omitempty"` 
	Item_recoil_max string  `xml:"item_recoil_max,attr,omitempty" json:"item_recoil_max,omitempty"` 
	Item_recoil_offset_coeff string  `xml:"item_recoil_offset_coeff,attr,omitempty" json:"item_recoil_offset_coeff,omitempty"` 
	Item_recoil_recovery_speed string  `xml:"item_recoil_recovery_speed,attr,omitempty" json:"item_recoil_recovery_speed,omitempty"` 
	Item_recoil_rotation_coeff string  `xml:"item_recoil_rotation_coeff,attr,omitempty" json:"item_recoil_rotation_coeff,omitempty"` 
	MChargeCount string  `xml:"mChargeCount,attr,omitempty" json:"mChargeCount,omitempty"` 
	MIsInitialized string  `xml:"mIsInitialized,attr,omitempty" json:"mIsInitialized,omitempty"` 
	Mana string  `xml:"mana,attr,omitempty" json:"mana,omitempty"` 
	Mana_charge_speed string  `xml:"mana_charge_speed,attr,omitempty" json:"mana_charge_speed,omitempty"` 
	Mana_max string  `xml:"mana_max,attr,omitempty" json:"mana_max,omitempty"` 
	Max_amount_in_inventory string  `xml:"max_amount_in_inventory,attr,omitempty" json:"max_amount_in_inventory,omitempty"` 
	Max_charged_actions string  `xml:"max_charged_actions,attr,omitempty" json:"max_charged_actions,omitempty"` 
	Never_reload string  `xml:"never_reload,attr,omitempty" json:"never_reload,omitempty"` 
	Reload_time_frames string  `xml:"reload_time_frames,attr,omitempty" json:"reload_time_frames,omitempty"` 
	Rotate_hand_amount string  `xml:"rotate_hand_amount,attr,omitempty" json:"rotate_hand_amount,omitempty"` 
	Rotate_in_hand string  `xml:"rotate_in_hand,attr,omitempty" json:"rotate_in_hand,omitempty"` 
	Rotate_in_hand_amount string  `xml:"rotate_in_hand_amount,attr,omitempty" json:"rotate_in_hand_amount,omitempty"` 
	Shooting_reduces_amount_in_inventory string  `xml:"shooting_reduces_amount_in_inventory,attr,omitempty" json:"shooting_reduces_amount_in_inventory,omitempty"` 
	Simulate_throw_as_item string  `xml:"simulate_throw_as_item,attr,omitempty" json:"simulate_throw_as_item,omitempty"` 
	Slot_consumption_function string  `xml:"slot_consumption_function,attr,omitempty" json:"slot_consumption_function,omitempty"` 
	Sprite_file string  `xml:"sprite_file,attr,omitempty" json:"sprite_file,omitempty"` 
	Stat_times_player_has_edited string  `xml:"stat_times_player_has_edited,attr,omitempty" json:"stat_times_player_has_edited,omitempty"` 
	Stat_times_player_has_shot string  `xml:"stat_times_player_has_shot,attr,omitempty" json:"stat_times_player_has_shot,omitempty"` 
	Swim_propel_amount string  `xml:"swim_propel_amount,attr,omitempty" json:"swim_propel_amount,omitempty"` 
	Throw_as_item string  `xml:"throw_as_item,attr,omitempty" json:"throw_as_item,omitempty"` 
	Ui_name string  `xml:"ui_name,attr,omitempty" json:"ui_name,omitempty"` 
	Use_entity_file_as_projectile_info_proxy string  `xml:"use_entity_file_as_projectile_info_proxy,attr,omitempty" json:"use_entity_file_as_projectile_info_proxy,omitempty"` 
	Use_gun_script string  `xml:"use_gun_script,attr,omitempty" json:"use_gun_script,omitempty"` 
	Gun_config []Gun_config_Struct  `xml:"gun_config,omitempty" json:"gun_config,omitempty"` 
	Gunaction_config []Gunaction_config_Struct  `xml:"gunaction_config,omitempty" json:"gunaction_config,omitempty"` 
}

type Gun_config_Struct struct {
	Actions_per_round string  `xml:"actions_per_round,attr,omitempty" json:"actions_per_round,omitempty"` 
	Deck_capacity string  `xml:"deck_capacity,attr,omitempty" json:"deck_capacity,omitempty"` 
	Reload_time string  `xml:"reload_time,attr,omitempty" json:"reload_time,omitempty"` 
	Shuffle_deck_when_empty string  `xml:"shuffle_deck_when_empty,attr,omitempty" json:"shuffle_deck_when_empty,omitempty"` 
}

type Gunaction_config_Struct struct {
	Action_ai_never_uses string  `xml:"action_ai_never_uses,attr,omitempty" json:"action_ai_never_uses,omitempty"` 
	Action_draw_many_count string  `xml:"action_draw_many_count,attr,omitempty" json:"action_draw_many_count,omitempty"` 
	Action_is_dangerous_blast string  `xml:"action_is_dangerous_blast,attr,omitempty" json:"action_is_dangerous_blast,omitempty"` 
	Action_mana_drain string  `xml:"action_mana_drain,attr,omitempty" json:"action_mana_drain,omitempty"` 
	Action_max_uses string  `xml:"action_max_uses,attr,omitempty" json:"action_max_uses,omitempty"` 
	Action_never_unlimited string  `xml:"action_never_unlimited,attr,omitempty" json:"action_never_unlimited,omitempty"` 
	Action_spawn_manual_unlock string  `xml:"action_spawn_manual_unlock,attr,omitempty" json:"action_spawn_manual_unlock,omitempty"` 
	Action_type string  `xml:"action_type,attr,omitempty" json:"action_type,omitempty"` 
	Action_unidentified_sprite_filename string  `xml:"action_unidentified_sprite_filename,attr,omitempty" json:"action_unidentified_sprite_filename,omitempty"` 
	Blood_count_multiplier string  `xml:"blood_count_multiplier,attr,omitempty" json:"blood_count_multiplier,omitempty"` 
	Bounces string  `xml:"bounces,attr,omitempty" json:"bounces,omitempty"` 
	Child_speed_multiplier string  `xml:"child_speed_multiplier,attr,omitempty" json:"child_speed_multiplier,omitempty"` 
	Damage_critical_chance string  `xml:"damage_critical_chance,attr,omitempty" json:"damage_critical_chance,omitempty"` 
	Damage_critical_multiplier string  `xml:"damage_critical_multiplier,attr,omitempty" json:"damage_critical_multiplier,omitempty"` 
	Damage_curse_add string  `xml:"damage_curse_add,attr,omitempty" json:"damage_curse_add,omitempty"` 
	Damage_drill_add string  `xml:"damage_drill_add,attr,omitempty" json:"damage_drill_add,omitempty"` 
	Damage_electricity_add string  `xml:"damage_electricity_add,attr,omitempty" json:"damage_electricity_add,omitempty"` 
	Damage_explosion_add string  `xml:"damage_explosion_add,attr,omitempty" json:"damage_explosion_add,omitempty"` 
	Damage_fire_add string  `xml:"damage_fire_add,attr,omitempty" json:"damage_fire_add,omitempty"` 
	Damage_healing_add string  `xml:"damage_healing_add,attr,omitempty" json:"damage_healing_add,omitempty"` 
	Damage_ice_add string  `xml:"damage_ice_add,attr,omitempty" json:"damage_ice_add,omitempty"` 
	Damage_melee_add string  `xml:"damage_melee_add,attr,omitempty" json:"damage_melee_add,omitempty"` 
	Damage_projectile_add string  `xml:"damage_projectile_add,attr,omitempty" json:"damage_projectile_add,omitempty"` 
	Damage_slice_add string  `xml:"damage_slice_add,attr,omitempty" json:"damage_slice_add,omitempty"` 
	Dampening string  `xml:"dampening,attr,omitempty" json:"dampening,omitempty"` 
	Explosion_damage_to_materials string  `xml:"explosion_damage_to_materials,attr,omitempty" json:"explosion_damage_to_materials,omitempty"` 
	Explosion_radius string  `xml:"explosion_radius,attr,omitempty" json:"explosion_radius,omitempty"` 
	Fire_rate_wait string  `xml:"fire_rate_wait,attr,omitempty" json:"fire_rate_wait,omitempty"` 
	Friendly_fire string  `xml:"friendly_fire,attr,omitempty" json:"friendly_fire,omitempty"` 
	Gore_particles string  `xml:"gore_particles,attr,omitempty" json:"gore_particles,omitempty"` 
	Gravity string  `xml:"gravity,attr,omitempty" json:"gravity,omitempty"` 
	Knockback_force string  `xml:"knockback_force,attr,omitempty" json:"knockback_force,omitempty"` 
	Lifetime_add string  `xml:"lifetime_add,attr,omitempty" json:"lifetime_add,omitempty"` 
	Light string  `xml:"light,attr,omitempty" json:"light,omitempty"` 
	Lightning_count string  `xml:"lightning_count,attr,omitempty" json:"lightning_count,omitempty"` 
	Material_amount string  `xml:"material_amount,attr,omitempty" json:"material_amount,omitempty"` 
	Pattern_degrees string  `xml:"pattern_degrees,attr,omitempty" json:"pattern_degrees,omitempty"` 
	Physics_impulse_coeff string  `xml:"physics_impulse_coeff,attr,omitempty" json:"physics_impulse_coeff,omitempty"` 
	Ragdoll_fx string  `xml:"ragdoll_fx,attr,omitempty" json:"ragdoll_fx,omitempty"` 
	Recoil string  `xml:"recoil,attr,omitempty" json:"recoil,omitempty"` 
	Reload_time string  `xml:"reload_time,attr,omitempty" json:"reload_time,omitempty"` 
	Screenshake string  `xml:"screenshake,attr,omitempty" json:"screenshake,omitempty"` 
	Speed_multiplier string  `xml:"speed_multiplier,attr,omitempty" json:"speed_multiplier,omitempty"` 
	Spread_degrees string  `xml:"spread_degrees,attr,omitempty" json:"spread_degrees,omitempty"` 
	State_cards_drawn string  `xml:"state_cards_drawn,attr,omitempty" json:"state_cards_drawn,omitempty"` 
	State_destroyed_action string  `xml:"state_destroyed_action,attr,omitempty" json:"state_destroyed_action,omitempty"` 
	State_discarded_action string  `xml:"state_discarded_action,attr,omitempty" json:"state_discarded_action,omitempty"` 
	State_shuffled string  `xml:"state_shuffled,attr,omitempty" json:"state_shuffled,omitempty"` 
	Trail_material_amount string  `xml:"trail_material_amount,attr,omitempty" json:"trail_material_amount,omitempty"` 
}

type ItemAlchemyComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Material_animate_wand string  `xml:"material_animate_wand,attr,omitempty" json:"material_animate_wand,omitempty"` 
	Material_animate_wand_alt string  `xml:"material_animate_wand_alt,attr,omitempty" json:"material_animate_wand_alt,omitempty"` 
	Material_increase_uses_remaining string  `xml:"material_increase_uses_remaining,attr,omitempty" json:"material_increase_uses_remaining,omitempty"` 
	Material_make_always_cast string  `xml:"material_make_always_cast,attr,omitempty" json:"material_make_always_cast,omitempty"` 
	Material_remove_shuffle string  `xml:"material_remove_shuffle,attr,omitempty" json:"material_remove_shuffle,omitempty"` 
	Material_sacrifice string  `xml:"material_sacrifice,attr,omitempty" json:"material_sacrifice,omitempty"` 
}

type ItemComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Always_use_item_name_in_ui string  `xml:"always_use_item_name_in_ui,attr,omitempty" json:"always_use_item_name_in_ui,omitempty"` 
	Auto_pickup string  `xml:"auto_pickup,attr,omitempty" json:"auto_pickup,omitempty"` 
	Camera_max_distance string  `xml:"camera_max_distance,attr,omitempty" json:"camera_max_distance,omitempty"` 
	Camera_smooth_speed_multiplier string  `xml:"camera_smooth_speed_multiplier,attr,omitempty" json:"camera_smooth_speed_multiplier,omitempty"` 
	Collect_nondefault_actions string  `xml:"collect_nondefault_actions,attr,omitempty" json:"collect_nondefault_actions,omitempty"` 
	Drinkable string  `xml:"drinkable,attr,omitempty" json:"drinkable,omitempty"` 
	Enable_orb_hacks string  `xml:"enable_orb_hacks,attr,omitempty" json:"enable_orb_hacks,omitempty"` 
	Has_been_picked_by_player string  `xml:"has_been_picked_by_player,attr,omitempty" json:"has_been_picked_by_player,omitempty"` 
	Inventory_slot_x string  `xml:"inventory_slot.x,attr,omitempty" json:"inventory_slot.x,omitempty"` 
	Inventory_slot_y string  `xml:"inventory_slot.y,attr,omitempty" json:"inventory_slot.y,omitempty"` 
	Is_all_spells_book string  `xml:"is_all_spells_book,attr,omitempty" json:"is_all_spells_book,omitempty"` 
	Is_consumable string  `xml:"is_consumable,attr,omitempty" json:"is_consumable,omitempty"` 
	Is_equipable_forced string  `xml:"is_equipable_forced,attr,omitempty" json:"is_equipable_forced,omitempty"` 
	Is_frozen string  `xml:"is_frozen,attr,omitempty" json:"is_frozen,omitempty"` 
	Is_hittable_always string  `xml:"is_hittable_always,attr,omitempty" json:"is_hittable_always,omitempty"` 
	Is_identified string  `xml:"is_identified,attr,omitempty" json:"is_identified,omitempty"` 
	Is_pickable string  `xml:"is_pickable,attr,omitempty" json:"is_pickable,omitempty"` 
	Is_stackable string  `xml:"is_stackable,attr,omitempty" json:"is_stackable,omitempty"` 
	Item_pickup_radius string  `xml:"item_pickup_radius,attr,omitempty" json:"item_pickup_radius,omitempty"` 
	MFramePickedUp string  `xml:"mFramePickedUp,attr,omitempty" json:"mFramePickedUp,omitempty"` 
	Max_child_items string  `xml:"max_child_items,attr,omitempty" json:"max_child_items,omitempty"` 
	Next_frame_pickable string  `xml:"next_frame_pickable,attr,omitempty" json:"next_frame_pickable,omitempty"` 
	Npc_next_frame_pickable string  `xml:"npc_next_frame_pickable,attr,omitempty" json:"npc_next_frame_pickable,omitempty"` 
	Permanently_attached string  `xml:"permanently_attached,attr,omitempty" json:"permanently_attached,omitempty"` 
	Play_hover_animation string  `xml:"play_hover_animation,attr,omitempty" json:"play_hover_animation,omitempty"` 
	Play_pick_sound string  `xml:"play_pick_sound,attr,omitempty" json:"play_pick_sound,omitempty"` 
	Play_spinning_animation string  `xml:"play_spinning_animation,attr,omitempty" json:"play_spinning_animation,omitempty"` 
	Preferred_inventory string  `xml:"preferred_inventory,attr,omitempty" json:"preferred_inventory,omitempty"` 
	Remove_default_child_actions_on_death string  `xml:"remove_default_child_actions_on_death,attr,omitempty" json:"remove_default_child_actions_on_death,omitempty"` 
	Remove_on_death string  `xml:"remove_on_death,attr,omitempty" json:"remove_on_death,omitempty"` 
	Remove_on_death_if_empty string  `xml:"remove_on_death_if_empty,attr,omitempty" json:"remove_on_death_if_empty,omitempty"` 
	Spawn_pos_x string  `xml:"spawn_pos.x,attr,omitempty" json:"spawn_pos.x,omitempty"` 
	Spawn_pos_y string  `xml:"spawn_pos.y,attr,omitempty" json:"spawn_pos.y,omitempty"` 
	Stats_count_as_item_pick_up string  `xml:"stats_count_as_item_pick_up,attr,omitempty" json:"stats_count_as_item_pick_up,omitempty"` 
	Ui_display_description_on_pick_up_hint string  `xml:"ui_display_description_on_pick_up_hint,attr,omitempty" json:"ui_display_description_on_pick_up_hint,omitempty"` 
	Uses_remaining string  `xml:"uses_remaining,attr,omitempty" json:"uses_remaining,omitempty"` 
	Item_name string  `xml:"item_name,attr,omitempty" json:"item_name,omitempty"` 
	Ui_description string  `xml:"ui_description,attr,omitempty" json:"ui_description,omitempty"` 
	Ui_sprite string  `xml:"ui_sprite,attr,omitempty" json:"ui_sprite,omitempty"` 
}

type ManaReloaderComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
}

type MaterialAreaCheckerComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Area_aabb_max_x string  `xml:"area_aabb.max_x,attr,omitempty" json:"area_aabb.max_x,omitempty"` 
	Area_aabb_max_y string  `xml:"area_aabb.max_y,attr,omitempty" json:"area_aabb.max_y,omitempty"` 
	Area_aabb_min_x string  `xml:"area_aabb.min_x,attr,omitempty" json:"area_aabb.min_x,omitempty"` 
	Area_aabb_min_y string  `xml:"area_aabb.min_y,attr,omitempty" json:"area_aabb.min_y,omitempty"` 
	Kill_after_message string  `xml:"kill_after_message,attr,omitempty" json:"kill_after_message,omitempty"` 
	Look_for_failure string  `xml:"look_for_failure,attr,omitempty" json:"look_for_failure,omitempty"` 
	Material string  `xml:"material,attr,omitempty" json:"material,omitempty"` 
	Material2 string  `xml:"material2,attr,omitempty" json:"material2,omitempty"` 
	Update_every_x_frame string  `xml:"update_every_x_frame,attr,omitempty" json:"update_every_x_frame,omitempty"` 
}

type SimplePhysicsComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Can_go_up string  `xml:"can_go_up,attr,omitempty" json:"can_go_up,omitempty"` 
}

type ItemActionComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Action_id string  `xml:"action_id,attr,omitempty" json:"action_id,omitempty"` 
}

type SpriteOffsetAnimatorComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Sprite_id string  `xml:"sprite_id,attr,omitempty" json:"sprite_id,omitempty"` 
	X_amount string  `xml:"x_amount,attr,omitempty" json:"x_amount,omitempty"` 
	X_phase string  `xml:"x_phase,attr,omitempty" json:"x_phase,omitempty"` 
	X_phase_offset string  `xml:"x_phase_offset,attr,omitempty" json:"x_phase_offset,omitempty"` 
	X_speed string  `xml:"x_speed,attr,omitempty" json:"x_speed,omitempty"` 
	Y_amount string  `xml:"y_amount,attr,omitempty" json:"y_amount,omitempty"` 
	Y_speed string  `xml:"y_speed,attr,omitempty" json:"y_speed,omitempty"` 
}

type ElectricChargeComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Charge string  `xml:"charge,attr,omitempty" json:"charge,omitempty"` 
	Charge_time_frames string  `xml:"charge_time_frames,attr,omitempty" json:"charge_time_frames,omitempty"` 
	Electricity_emission_interval_frames string  `xml:"electricity_emission_interval_frames,attr,omitempty" json:"electricity_emission_interval_frames,omitempty"` 
	Fx_emission_interval_max string  `xml:"fx_emission_interval_max,attr,omitempty" json:"fx_emission_interval_max,omitempty"` 
	Fx_emission_interval_min string  `xml:"fx_emission_interval_min,attr,omitempty" json:"fx_emission_interval_min,omitempty"` 
	Fx_velocity_max string  `xml:"fx_velocity_max,attr,omitempty" json:"fx_velocity_max,omitempty"` 
}

type ItemAIKnowledgeComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Is_consumed string  `xml:"is_consumed,attr,omitempty" json:"is_consumed,omitempty"` 
	Is_known string  `xml:"is_known,attr,omitempty" json:"is_known,omitempty"` 
	Is_melee_weapon string  `xml:"is_melee_weapon,attr,omitempty" json:"is_melee_weapon,omitempty"` 
	Is_other_buffing string  `xml:"is_other_buffing,attr,omitempty" json:"is_other_buffing,omitempty"` 
	Is_other_healing string  `xml:"is_other_healing,attr,omitempty" json:"is_other_healing,omitempty"` 
	Is_ranged_weapon string  `xml:"is_ranged_weapon,attr,omitempty" json:"is_ranged_weapon,omitempty"` 
	Is_safe string  `xml:"is_safe,attr,omitempty" json:"is_safe,omitempty"` 
	Is_self_buffing string  `xml:"is_self_buffing,attr,omitempty" json:"is_self_buffing,omitempty"` 
	Is_self_healing string  `xml:"is_self_healing,attr,omitempty" json:"is_self_healing,omitempty"` 
	Is_throwable_weapon string  `xml:"is_throwable_weapon,attr,omitempty" json:"is_throwable_weapon,omitempty"` 
	Is_weapon string  `xml:"is_weapon,attr,omitempty" json:"is_weapon,omitempty"` 
	Never_use string  `xml:"never_use,attr,omitempty" json:"never_use,omitempty"` 
	Ranged_min_distance string  `xml:"ranged_min_distance,attr,omitempty" json:"ranged_min_distance,omitempty"` 
}

type UIInfoComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Name string  `xml:"name,attr,omitempty" json:"name,omitempty"` 
}

type ExplodeOnDamageComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Explode_on_damage_percent string  `xml:"explode_on_damage_percent,attr,omitempty" json:"explode_on_damage_percent,omitempty"` 
	Explode_on_death_percent string  `xml:"explode_on_death_percent,attr,omitempty" json:"explode_on_death_percent,omitempty"` 
	Physics_body_destruction_required string  `xml:"physics_body_destruction_required,attr,omitempty" json:"physics_body_destruction_required,omitempty"` 
	Physics_body_modified_death_probability string  `xml:"physics_body_modified_death_probability,attr,omitempty" json:"physics_body_modified_death_probability,omitempty"` 
	Config_explosion []Config_explosion_Struct  `xml:"config_explosion,omitempty" json:"config_explosion,omitempty"` 
}

type Damage_critical_Struct struct {
	Chance string  `xml:"chance,attr,omitempty" json:"chance,omitempty"` 
	Damage_multiplier string  `xml:"damage_multiplier,attr,omitempty" json:"damage_multiplier,omitempty"` 
}

type PhysicsBodyComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Allow_sleep string  `xml:"allow_sleep,attr,omitempty" json:"allow_sleep,omitempty"` 
	Angular_damping string  `xml:"angular_damping,attr,omitempty" json:"angular_damping,omitempty"` 
	Auto_clean string  `xml:"auto_clean,attr,omitempty" json:"auto_clean,omitempty"` 
	Buoyancy string  `xml:"buoyancy,attr,omitempty" json:"buoyancy,omitempty"` 
	Fixed_rotation string  `xml:"fixed_rotation,attr,omitempty" json:"fixed_rotation,omitempty"` 
	Force_add_update_areas string  `xml:"force_add_update_areas,attr,omitempty" json:"force_add_update_areas,omitempty"` 
	Go_through_sand string  `xml:"go_through_sand,attr,omitempty" json:"go_through_sand,omitempty"` 
	Gravity_scale_if_has_no_image_shapes string  `xml:"gravity_scale_if_has_no_image_shapes,attr,omitempty" json:"gravity_scale_if_has_no_image_shapes,omitempty"` 
	Gridworld_box2d string  `xml:"gridworld_box2d,attr,omitempty" json:"gridworld_box2d,omitempty"` 
	Hax_fix_going_through_ground string  `xml:"hax_fix_going_through_ground,attr,omitempty" json:"hax_fix_going_through_ground,omitempty"` 
	Hax_fix_going_through_sand string  `xml:"hax_fix_going_through_sand,attr,omitempty" json:"hax_fix_going_through_sand,omitempty"` 
	Hax_wait_till_pixel_scenes_loaded string  `xml:"hax_wait_till_pixel_scenes_loaded,attr,omitempty" json:"hax_wait_till_pixel_scenes_loaded,omitempty"` 
	Initial_velocity_x string  `xml:"initial_velocity.x,attr,omitempty" json:"initial_velocity.x,omitempty"` 
	Initial_velocity_y string  `xml:"initial_velocity.y,attr,omitempty" json:"initial_velocity.y,omitempty"` 
	Is_bullet string  `xml:"is_bullet,attr,omitempty" json:"is_bullet,omitempty"` 
	Is_character string  `xml:"is_character,attr,omitempty" json:"is_character,omitempty"` 
	Is_enabled string  `xml:"is_enabled,attr,omitempty" json:"is_enabled,omitempty"` 
	Is_external string  `xml:"is_external,attr,omitempty" json:"is_external,omitempty"` 
	Is_kinematic string  `xml:"is_kinematic,attr,omitempty" json:"is_kinematic,omitempty"` 
	Is_static string  `xml:"is_static,attr,omitempty" json:"is_static,omitempty"` 
	Kills_entity string  `xml:"kills_entity,attr,omitempty" json:"kills_entity,omitempty"` 
	Linear_damping string  `xml:"linear_damping,attr,omitempty" json:"linear_damping,omitempty"` 
	MActiveState string  `xml:"mActiveState,attr,omitempty" json:"mActiveState,omitempty"` 
	On_death_leave_physics_body string  `xml:"on_death_leave_physics_body,attr,omitempty" json:"on_death_leave_physics_body,omitempty"` 
	On_death_really_leave_body string  `xml:"on_death_really_leave_body,attr,omitempty" json:"on_death_really_leave_body,omitempty"` 
	Projectiles_rotate_toward_velocity string  `xml:"projectiles_rotate_toward_velocity,attr,omitempty" json:"projectiles_rotate_toward_velocity,omitempty"` 
	Randomize_init_velocity string  `xml:"randomize_init_velocity,attr,omitempty" json:"randomize_init_velocity,omitempty"` 
	Uid string  `xml:"uid,attr,omitempty" json:"uid,omitempty"` 
	Update_entity_transform string  `xml:"update_entity_transform,attr,omitempty" json:"update_entity_transform,omitempty"` 
}

type PhysicsImageShapeComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Body_id string  `xml:"body_id,attr,omitempty" json:"body_id,omitempty"` 
	Centered string  `xml:"centered,attr,omitempty" json:"centered,omitempty"` 
	Image_file string  `xml:"image_file,attr,omitempty" json:"image_file,omitempty"` 
	Is_circle string  `xml:"is_circle,attr,omitempty" json:"is_circle,omitempty"` 
	Is_root string  `xml:"is_root,attr,omitempty" json:"is_root,omitempty"` 
	Material string  `xml:"material,attr,omitempty" json:"material,omitempty"` 
	Offset_x string  `xml:"offset_x,attr,omitempty" json:"offset_x,omitempty"` 
	Offset_y string  `xml:"offset_y,attr,omitempty" json:"offset_y,omitempty"` 
	Use_sprite string  `xml:"use_sprite,attr,omitempty" json:"use_sprite,omitempty"` 
	Z string  `xml:"z,attr,omitempty" json:"z,omitempty"` 
}

type PhysicsThrowableComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Attach_min_speed string  `xml:"attach_min_speed,attr,omitempty" json:"attach_min_speed,omitempty"` 
	Attach_to_surfaces_knife_style string  `xml:"attach_to_surfaces_knife_style,attr,omitempty" json:"attach_to_surfaces_knife_style,omitempty"` 
	Hp string  `xml:"hp,attr,omitempty" json:"hp,omitempty"` 
	Max_throw_speed string  `xml:"max_throw_speed,attr,omitempty" json:"max_throw_speed,omitempty"` 
	Max_torque string  `xml:"max_torque,attr,omitempty" json:"max_torque,omitempty"` 
	Min_torque string  `xml:"min_torque,attr,omitempty" json:"min_torque,omitempty"` 
	Throw_force_coeff string  `xml:"throw_force_coeff,attr,omitempty" json:"throw_force_coeff,omitempty"` 
	Tip_check_offset_max string  `xml:"tip_check_offset_max,attr,omitempty" json:"tip_check_offset_max,omitempty"` 
	Tip_check_offset_min string  `xml:"tip_check_offset_min,attr,omitempty" json:"tip_check_offset_min,omitempty"` 
	Tip_check_random_rotation_deg string  `xml:"tip_check_random_rotation_deg,attr,omitempty" json:"tip_check_random_rotation_deg,omitempty"` 
}

type PotionComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Body_colored string  `xml:"body_colored,attr,omitempty" json:"body_colored,omitempty"` 
	Never_color string  `xml:"never_color,attr,omitempty" json:"never_color,omitempty"` 
	Spray_velocity_coeff string  `xml:"spray_velocity_coeff,attr,omitempty" json:"spray_velocity_coeff,omitempty"` 
	Spray_velocity_normalized_min string  `xml:"spray_velocity_normalized_min,attr,omitempty" json:"spray_velocity_normalized_min,omitempty"` 
	Throw_bunch string  `xml:"throw_bunch,attr,omitempty" json:"throw_bunch,omitempty"` 
	Throw_how_many string  `xml:"throw_how_many,attr,omitempty" json:"throw_how_many,omitempty"` 
}

type Config_Struct struct {
	Action_ai_never_uses string  `xml:"action_ai_never_uses,attr,omitempty" json:"action_ai_never_uses,omitempty"` 
	Action_draw_many_count string  `xml:"action_draw_many_count,attr,omitempty" json:"action_draw_many_count,omitempty"` 
	Action_is_dangerous_blast string  `xml:"action_is_dangerous_blast,attr,omitempty" json:"action_is_dangerous_blast,omitempty"` 
	Action_mana_drain string  `xml:"action_mana_drain,attr,omitempty" json:"action_mana_drain,omitempty"` 
	Action_max_uses string  `xml:"action_max_uses,attr,omitempty" json:"action_max_uses,omitempty"` 
	Action_never_unlimited string  `xml:"action_never_unlimited,attr,omitempty" json:"action_never_unlimited,omitempty"` 
	Action_spawn_manual_unlock string  `xml:"action_spawn_manual_unlock,attr,omitempty" json:"action_spawn_manual_unlock,omitempty"` 
	Action_type string  `xml:"action_type,attr,omitempty" json:"action_type,omitempty"` 
	Action_unidentified_sprite_filename string  `xml:"action_unidentified_sprite_filename,attr,omitempty" json:"action_unidentified_sprite_filename,omitempty"` 
	Blood_count_multiplier string  `xml:"blood_count_multiplier,attr,omitempty" json:"blood_count_multiplier,omitempty"` 
	Bounces string  `xml:"bounces,attr,omitempty" json:"bounces,omitempty"` 
	Child_speed_multiplier string  `xml:"child_speed_multiplier,attr,omitempty" json:"child_speed_multiplier,omitempty"` 
	Damage_critical_chance string  `xml:"damage_critical_chance,attr,omitempty" json:"damage_critical_chance,omitempty"` 
	Damage_critical_multiplier string  `xml:"damage_critical_multiplier,attr,omitempty" json:"damage_critical_multiplier,omitempty"` 
	Damage_curse_add string  `xml:"damage_curse_add,attr,omitempty" json:"damage_curse_add,omitempty"` 
	Damage_drill_add string  `xml:"damage_drill_add,attr,omitempty" json:"damage_drill_add,omitempty"` 
	Damage_electricity_add string  `xml:"damage_electricity_add,attr,omitempty" json:"damage_electricity_add,omitempty"` 
	Damage_explosion_add string  `xml:"damage_explosion_add,attr,omitempty" json:"damage_explosion_add,omitempty"` 
	Damage_fire_add string  `xml:"damage_fire_add,attr,omitempty" json:"damage_fire_add,omitempty"` 
	Damage_healing_add string  `xml:"damage_healing_add,attr,omitempty" json:"damage_healing_add,omitempty"` 
	Damage_ice_add string  `xml:"damage_ice_add,attr,omitempty" json:"damage_ice_add,omitempty"` 
	Damage_melee_add string  `xml:"damage_melee_add,attr,omitempty" json:"damage_melee_add,omitempty"` 
	Damage_projectile_add string  `xml:"damage_projectile_add,attr,omitempty" json:"damage_projectile_add,omitempty"` 
	Damage_slice_add string  `xml:"damage_slice_add,attr,omitempty" json:"damage_slice_add,omitempty"` 
	Dampening string  `xml:"dampening,attr,omitempty" json:"dampening,omitempty"` 
	Explosion_damage_to_materials string  `xml:"explosion_damage_to_materials,attr,omitempty" json:"explosion_damage_to_materials,omitempty"` 
	Explosion_radius string  `xml:"explosion_radius,attr,omitempty" json:"explosion_radius,omitempty"` 
	Fire_rate_wait string  `xml:"fire_rate_wait,attr,omitempty" json:"fire_rate_wait,omitempty"` 
	Friendly_fire string  `xml:"friendly_fire,attr,omitempty" json:"friendly_fire,omitempty"` 
	Gore_particles string  `xml:"gore_particles,attr,omitempty" json:"gore_particles,omitempty"` 
	Gravity string  `xml:"gravity,attr,omitempty" json:"gravity,omitempty"` 
	Knockback_force string  `xml:"knockback_force,attr,omitempty" json:"knockback_force,omitempty"` 
	Lifetime_add string  `xml:"lifetime_add,attr,omitempty" json:"lifetime_add,omitempty"` 
	Light string  `xml:"light,attr,omitempty" json:"light,omitempty"` 
	Lightning_count string  `xml:"lightning_count,attr,omitempty" json:"lightning_count,omitempty"` 
	Material_amount string  `xml:"material_amount,attr,omitempty" json:"material_amount,omitempty"` 
	Pattern_degrees string  `xml:"pattern_degrees,attr,omitempty" json:"pattern_degrees,omitempty"` 
	Physics_impulse_coeff string  `xml:"physics_impulse_coeff,attr,omitempty" json:"physics_impulse_coeff,omitempty"` 
	Ragdoll_fx string  `xml:"ragdoll_fx,attr,omitempty" json:"ragdoll_fx,omitempty"` 
	Recoil string  `xml:"recoil,attr,omitempty" json:"recoil,omitempty"` 
	Reload_time string  `xml:"reload_time,attr,omitempty" json:"reload_time,omitempty"` 
	Screenshake string  `xml:"screenshake,attr,omitempty" json:"screenshake,omitempty"` 
	Speed_multiplier string  `xml:"speed_multiplier,attr,omitempty" json:"speed_multiplier,omitempty"` 
	Spread_degrees string  `xml:"spread_degrees,attr,omitempty" json:"spread_degrees,omitempty"` 
	State_cards_drawn string  `xml:"state_cards_drawn,attr,omitempty" json:"state_cards_drawn,omitempty"` 
	State_destroyed_action string  `xml:"state_destroyed_action,attr,omitempty" json:"state_destroyed_action,omitempty"` 
	State_discarded_action string  `xml:"state_discarded_action,attr,omitempty" json:"state_discarded_action,omitempty"` 
	State_shuffled string  `xml:"state_shuffled,attr,omitempty" json:"state_shuffled,omitempty"` 
	Trail_material_amount string  `xml:"trail_material_amount,attr,omitempty" json:"trail_material_amount,omitempty"` 
}

type Damage_by_type_Struct struct {
	Curse string  `xml:"curse,attr,omitempty" json:"curse,omitempty"` 
	Drill string  `xml:"drill,attr,omitempty" json:"drill,omitempty"` 
	Electricity string  `xml:"electricity,attr,omitempty" json:"electricity,omitempty"` 
	Explosion string  `xml:"explosion,attr,omitempty" json:"explosion,omitempty"` 
	Fire string  `xml:"fire,attr,omitempty" json:"fire,omitempty"` 
	Healing string  `xml:"healing,attr,omitempty" json:"healing,omitempty"` 
	Ice string  `xml:"ice,attr,omitempty" json:"ice,omitempty"` 
	Melee string  `xml:"melee,attr,omitempty" json:"melee,omitempty"` 
	Overeating string  `xml:"overeating,attr,omitempty" json:"overeating,omitempty"` 
	Physics_hit string  `xml:"physics_hit,attr,omitempty" json:"physics_hit,omitempty"` 
	Poison string  `xml:"poison,attr,omitempty" json:"poison,omitempty"` 
	Projectile string  `xml:"projectile,attr,omitempty" json:"projectile,omitempty"` 
	Radioactive string  `xml:"radioactive,attr,omitempty" json:"radioactive,omitempty"` 
	Slice string  `xml:"slice,attr,omitempty" json:"slice,omitempty"` 
}

type PhysicsBodyCollisionDamageComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Damage_multiplier string  `xml:"damage_multiplier,attr,omitempty" json:"damage_multiplier,omitempty"` 
	Speed_threshold string  `xml:"speed_threshold,attr,omitempty" json:"speed_threshold,omitempty"` 
}

type GameEffectComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Caused_by_ingestion_status_effect string  `xml:"caused_by_ingestion_status_effect,attr,omitempty" json:"caused_by_ingestion_status_effect,omitempty"` 
	Causing_status_effect string  `xml:"causing_status_effect,attr,omitempty" json:"causing_status_effect,omitempty"` 
	Disable_movement string  `xml:"disable_movement,attr,omitempty" json:"disable_movement,omitempty"` 
	Effect string  `xml:"effect,attr,omitempty" json:"effect,omitempty"` 
	Exclusivity_group string  `xml:"exclusivity_group,attr,omitempty" json:"exclusivity_group,omitempty"` 
	Frames string  `xml:"frames,attr,omitempty" json:"frames,omitempty"` 
	MCaster string  `xml:"mCaster,attr,omitempty" json:"mCaster,omitempty"` 
	MCasterHerdId string  `xml:"mCasterHerdId,attr,omitempty" json:"mCasterHerdId,omitempty"` 
	MCharmDisabledCameraBound string  `xml:"mCharmDisabledCameraBound,attr,omitempty" json:"mCharmDisabledCameraBound,omitempty"` 
	MCharmEnabledTeleporting string  `xml:"mCharmEnabledTeleporting,attr,omitempty" json:"mCharmEnabledTeleporting,omitempty"` 
	MCooldown string  `xml:"mCooldown,attr,omitempty" json:"mCooldown,omitempty"` 
	MCounter string  `xml:"mCounter,attr,omitempty" json:"mCounter,omitempty"` 
	MInvisible string  `xml:"mInvisible,attr,omitempty" json:"mInvisible,omitempty"` 
	MIsExtension string  `xml:"mIsExtension,attr,omitempty" json:"mIsExtension,omitempty"` 
	Ragdoll_effect string  `xml:"ragdoll_effect,attr,omitempty" json:"ragdoll_effect,omitempty"` 
	Ragdoll_fx_custom_entity_apply_only_to_largest_body string  `xml:"ragdoll_fx_custom_entity_apply_only_to_largest_body,attr,omitempty" json:"ragdoll_fx_custom_entity_apply_only_to_largest_body,omitempty"` 
	Ragdoll_material string  `xml:"ragdoll_material,attr,omitempty" json:"ragdoll_material,omitempty"` 
	Report_block_msg string  `xml:"report_block_msg,attr,omitempty" json:"report_block_msg,omitempty"` 
	Teleportation_delay_min_frames string  `xml:"teleportation_delay_min_frames,attr,omitempty" json:"teleportation_delay_min_frames,omitempty"` 
	Teleportation_probability string  `xml:"teleportation_probability,attr,omitempty" json:"teleportation_probability,omitempty"` 
	Teleportation_radius_max string  `xml:"teleportation_radius_max,attr,omitempty" json:"teleportation_radius_max,omitempty"` 
	Teleportation_radius_min string  `xml:"teleportation_radius_min,attr,omitempty" json:"teleportation_radius_min,omitempty"` 
	Teleportations_num string  `xml:"teleportations_num,attr,omitempty" json:"teleportations_num,omitempty"` 
}

type UIIconComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Description string  `xml:"description,attr,omitempty" json:"description,omitempty"` 
	Display_above_head string  `xml:"display_above_head,attr,omitempty" json:"display_above_head,omitempty"` 
	Display_in_hud string  `xml:"display_in_hud,attr,omitempty" json:"display_in_hud,omitempty"` 
	Icon_sprite_file string  `xml:"icon_sprite_file,attr,omitempty" json:"icon_sprite_file,omitempty"` 
	Is_perk string  `xml:"is_perk,attr,omitempty" json:"is_perk,omitempty"` 
	Name string  `xml:"name,attr,omitempty" json:"name,omitempty"` 
}

type MagicConvertMaterialComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Clean_stains string  `xml:"clean_stains,attr,omitempty" json:"clean_stains,omitempty"` 
	Convert_entities string  `xml:"convert_entities,attr,omitempty" json:"convert_entities,omitempty"` 
	Convert_same_material string  `xml:"convert_same_material,attr,omitempty" json:"convert_same_material,omitempty"` 
	Extinguish_fire string  `xml:"extinguish_fire,attr,omitempty" json:"extinguish_fire,omitempty"` 
	Fan_the_flames string  `xml:"fan_the_flames,attr,omitempty" json:"fan_the_flames,omitempty"` 
	From_any_material string  `xml:"from_any_material,attr,omitempty" json:"from_any_material,omitempty"` 
	From_material string  `xml:"from_material,attr,omitempty" json:"from_material,omitempty"` 
	Ignite_materials string  `xml:"ignite_materials,attr,omitempty" json:"ignite_materials,omitempty"` 
	Is_circle string  `xml:"is_circle,attr,omitempty" json:"is_circle,omitempty"` 
	Kill_when_finished string  `xml:"kill_when_finished,attr,omitempty" json:"kill_when_finished,omitempty"` 
	Loop string  `xml:"loop,attr,omitempty" json:"loop,omitempty"` 
	MRadius string  `xml:"mRadius,attr,omitempty" json:"mRadius,omitempty"` 
	Min_radius string  `xml:"min_radius,attr,omitempty" json:"min_radius,omitempty"` 
	Radius string  `xml:"radius,attr,omitempty" json:"radius,omitempty"` 
	Reaction_audio_amount string  `xml:"reaction_audio_amount,attr,omitempty" json:"reaction_audio_amount,omitempty"` 
	Stain_frozen string  `xml:"stain_frozen,attr,omitempty" json:"stain_frozen,omitempty"` 
	Steps_per_frame string  `xml:"steps_per_frame,attr,omitempty" json:"steps_per_frame,omitempty"` 
	Temperature_reaction_temp string  `xml:"temperature_reaction_temp,attr,omitempty" json:"temperature_reaction_temp,omitempty"` 
	To_material string  `xml:"to_material,attr,omitempty" json:"to_material,omitempty"` 
	From_material_tag string  `xml:"from_material_tag,attr,omitempty" json:"from_material_tag,omitempty"` 
}

type EnergyShieldComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Tags string  `xml:"_tags,attr,omitempty" json:"_tags,omitempty"` 
	Damage_multiplier string  `xml:"damage_multiplier,attr,omitempty" json:"damage_multiplier,omitempty"` 
	Energy string  `xml:"energy,attr,omitempty" json:"energy,omitempty"` 
	Energy_required_to_shield string  `xml:"energy_required_to_shield,attr,omitempty" json:"energy_required_to_shield,omitempty"` 
	Max_energy string  `xml:"max_energy,attr,omitempty" json:"max_energy,omitempty"` 
	Radius string  `xml:"radius,attr,omitempty" json:"radius,omitempty"` 
	Recharge_speed string  `xml:"recharge_speed,attr,omitempty" json:"recharge_speed,omitempty"` 
	Sector_degrees string  `xml:"sector_degrees,attr,omitempty" json:"sector_degrees,omitempty"` 
}

type LaserEmitterComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Emit_until_frame string  `xml:"emit_until_frame,attr,omitempty" json:"emit_until_frame,omitempty"` 
	Is_emitting string  `xml:"is_emitting,attr,omitempty" json:"is_emitting,omitempty"` 
	Laser_angle_add_rad string  `xml:"laser_angle_add_rad,attr,omitempty" json:"laser_angle_add_rad,omitempty"` 
	Laser []Laser_Struct  `xml:"laser,omitempty" json:"laser,omitempty"` 
}

type Laser_Struct struct {
	Audio_enabled string  `xml:"audio_enabled,attr,omitempty" json:"audio_enabled,omitempty"` 
	Audio_hit_always_enabled string  `xml:"audio_hit_always_enabled,attr,omitempty" json:"audio_hit_always_enabled,omitempty"` 
	Beam_particle_chance string  `xml:"beam_particle_chance,attr,omitempty" json:"beam_particle_chance,omitempty"` 
	Beam_particle_fade string  `xml:"beam_particle_fade,attr,omitempty" json:"beam_particle_fade,omitempty"` 
	Beam_particle_fade_reverse string  `xml:"beam_particle_fade_reverse,attr,omitempty" json:"beam_particle_fade_reverse,omitempty"` 
	Beam_particle_type string  `xml:"beam_particle_type,attr,omitempty" json:"beam_particle_type,omitempty"` 
	Beam_radius string  `xml:"beam_radius,attr,omitempty" json:"beam_radius,omitempty"` 
	Damage_to_cells string  `xml:"damage_to_cells,attr,omitempty" json:"damage_to_cells,omitempty"` 
	Damage_to_entities string  `xml:"damage_to_entities,attr,omitempty" json:"damage_to_entities,omitempty"` 
	Hit_particle_chance string  `xml:"hit_particle_chance,attr,omitempty" json:"hit_particle_chance,omitempty"` 
	Max_cell_durability_to_destroy string  `xml:"max_cell_durability_to_destroy,attr,omitempty" json:"max_cell_durability_to_destroy,omitempty"` 
	Max_length string  `xml:"max_length,attr,omitempty" json:"max_length,omitempty"` 
	Root_entity_is_responsible_for_damage string  `xml:"root_entity_is_responsible_for_damage,attr,omitempty" json:"root_entity_is_responsible_for_damage,omitempty"` 
}

type CellEaterComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Eat_dynamic_physics_bodies string  `xml:"eat_dynamic_physics_bodies,attr,omitempty" json:"eat_dynamic_physics_bodies,omitempty"` 
	Eat_probability string  `xml:"eat_probability,attr,omitempty" json:"eat_probability,omitempty"` 
	Ignored_material string  `xml:"ignored_material,attr,omitempty" json:"ignored_material,omitempty"` 
	Limited_materials string  `xml:"limited_materials,attr,omitempty" json:"limited_materials,omitempty"` 
	Materials string  `xml:"materials,attr,omitempty" json:"materials,omitempty"` 
	Only_stain string  `xml:"only_stain,attr,omitempty" json:"only_stain,omitempty"` 
	Radius string  `xml:"radius,attr,omitempty" json:"radius,omitempty"` 
}

type WorldStateComponent_Struct struct {
	DEBUG_LOADED_FROM_AUTOSAVE string  `xml:"DEBUG_LOADED_FROM_AUTOSAVE,attr,omitempty" json:"DEBUG_LOADED_FROM_AUTOSAVE,omitempty"` 
	DEBUG_LOADED_FROM_OLD_VERSION string  `xml:"DEBUG_LOADED_FROM_OLD_VERSION,attr,omitempty" json:"DEBUG_LOADED_FROM_OLD_VERSION,omitempty"` 
	ENDING_HAPPINESS string  `xml:"ENDING_HAPPINESS,attr,omitempty" json:"ENDING_HAPPINESS,omitempty"` 
	ENDING_HAPPINESS_FRAMES string  `xml:"ENDING_HAPPINESS_FRAMES,attr,omitempty" json:"ENDING_HAPPINESS_FRAMES,omitempty"` 
	ENDING_HAPPINESS_HAPPENING string  `xml:"ENDING_HAPPINESS_HAPPENING,attr,omitempty" json:"ENDING_HAPPINESS_HAPPENING,omitempty"` 
	EVERYTHING_TO_GOLD string  `xml:"EVERYTHING_TO_GOLD,attr,omitempty" json:"EVERYTHING_TO_GOLD,omitempty"` 
	INFINITE_GOLD_HAPPENING string  `xml:"INFINITE_GOLD_HAPPENING,attr,omitempty" json:"INFINITE_GOLD_HAPPENING,omitempty"` 
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Clouds_01_target string  `xml:"clouds_01_target,attr,omitempty" json:"clouds_01_target,omitempty"` 
	Clouds_02_target string  `xml:"clouds_02_target,attr,omitempty" json:"clouds_02_target,omitempty"` 
	Consume_actions string  `xml:"consume_actions,attr,omitempty" json:"consume_actions,omitempty"` 
	Damage_flash_multiplier string  `xml:"damage_flash_multiplier,attr,omitempty" json:"damage_flash_multiplier,omitempty"` 
	Day_count string  `xml:"day_count,attr,omitempty" json:"day_count,omitempty"` 
	Fog string  `xml:"fog,attr,omitempty" json:"fog,omitempty"` 
	Fog_target string  `xml:"fog_target,attr,omitempty" json:"fog_target,omitempty"` 
	Global_genome_relations_modifier string  `xml:"global_genome_relations_modifier,attr,omitempty" json:"global_genome_relations_modifier,omitempty"` 
	Gore_multiplier string  `xml:"gore_multiplier,attr,omitempty" json:"gore_multiplier,omitempty"` 
	Gradient_sky_alpha_target string  `xml:"gradient_sky_alpha_target,attr,omitempty" json:"gradient_sky_alpha_target,omitempty"` 
	Intro_weather string  `xml:"intro_weather,attr,omitempty" json:"intro_weather,omitempty"` 
	Is_initialized string  `xml:"is_initialized,attr,omitempty" json:"is_initialized,omitempty"` 
	Lightning_count string  `xml:"lightning_count,attr,omitempty" json:"lightning_count,omitempty"` 
	MFlashAlpha string  `xml:"mFlashAlpha,attr,omitempty" json:"mFlashAlpha,omitempty"` 
	Material_everything_to_gold string  `xml:"material_everything_to_gold,attr,omitempty" json:"material_everything_to_gold,omitempty"` 
	Material_everything_to_gold_static string  `xml:"material_everything_to_gold_static,attr,omitempty" json:"material_everything_to_gold_static,omitempty"` 
	Mods_have_been_active_during_this_run string  `xml:"mods_have_been_active_during_this_run,attr,omitempty" json:"mods_have_been_active_during_this_run,omitempty"` 
	Next_cut_through_world_id string  `xml:"next_cut_through_world_id,attr,omitempty" json:"next_cut_through_world_id,omitempty"` 
	Next_portal_id string  `xml:"next_portal_id,attr,omitempty" json:"next_portal_id,omitempty"` 
	Open_fog_of_war_everywhere string  `xml:"open_fog_of_war_everywhere,attr,omitempty" json:"open_fog_of_war_everywhere,omitempty"` 
	Perk_gold_is_forever string  `xml:"perk_gold_is_forever,attr,omitempty" json:"perk_gold_is_forever,omitempty"` 
	Perk_hp_drop_chance string  `xml:"perk_hp_drop_chance,attr,omitempty" json:"perk_hp_drop_chance,omitempty"` 
	Perk_infinite_spells string  `xml:"perk_infinite_spells,attr,omitempty" json:"perk_infinite_spells,omitempty"` 
	Perk_rats_player_friendly string  `xml:"perk_rats_player_friendly,attr,omitempty" json:"perk_rats_player_friendly,omitempty"` 
	Perk_trick_kills_blood_money string  `xml:"perk_trick_kills_blood_money,attr,omitempty" json:"perk_trick_kills_blood_money,omitempty"` 
	Player_did_damage_over_1milj string  `xml:"player_did_damage_over_1milj,attr,omitempty" json:"player_did_damage_over_1milj,omitempty"` 
	Player_did_infinite_spell_count string  `xml:"player_did_infinite_spell_count,attr,omitempty" json:"player_did_infinite_spell_count,omitempty"` 
	Player_living_with_minus_hp string  `xml:"player_living_with_minus_hp,attr,omitempty" json:"player_living_with_minus_hp,omitempty"` 
	Player_polymorph_count string  `xml:"player_polymorph_count,attr,omitempty" json:"player_polymorph_count,omitempty"` 
	Player_polymorph_random_count string  `xml:"player_polymorph_random_count,attr,omitempty" json:"player_polymorph_random_count,omitempty"` 
	Player_spawn_location_x string  `xml:"player_spawn_location.x,attr,omitempty" json:"player_spawn_location.x,omitempty"` 
	Player_spawn_location_y string  `xml:"player_spawn_location.y,attr,omitempty" json:"player_spawn_location.y,omitempty"` 
	Rain string  `xml:"rain,attr,omitempty" json:"rain,omitempty"` 
	Rain_target string  `xml:"rain_target,attr,omitempty" json:"rain_target,omitempty"` 
	Session_stat_file string  `xml:"session_stat_file,attr,omitempty" json:"session_stat_file,omitempty"` 
	Sky_sunset_alpha_target string  `xml:"sky_sunset_alpha_target,attr,omitempty" json:"sky_sunset_alpha_target,omitempty"` 
	Time string  `xml:"time,attr,omitempty" json:"time,omitempty"` 
	Time_dt string  `xml:"time_dt,attr,omitempty" json:"time_dt,omitempty"` 
	Time_total string  `xml:"time_total,attr,omitempty" json:"time_total,omitempty"` 
	Trick_kill_gold_multiplier string  `xml:"trick_kill_gold_multiplier,attr,omitempty" json:"trick_kill_gold_multiplier,omitempty"` 
	Twitch_has_been_active_during_this_run string  `xml:"twitch_has_been_active_during_this_run,attr,omitempty" json:"twitch_has_been_active_during_this_run,omitempty"` 
	Wind string  `xml:"wind,attr,omitempty" json:"wind,omitempty"` 
	Wind_speed string  `xml:"wind_speed,attr,omitempty" json:"wind_speed,omitempty"` 
	Wind_speed_sin string  `xml:"wind_speed_sin,attr,omitempty" json:"wind_speed_sin,omitempty"` 
	Wind_speed_sin_t string  `xml:"wind_speed_sin_t,attr,omitempty" json:"wind_speed_sin_t,omitempty"` 
	Lua_globals []Lua_globals_Struct  `xml:"lua_globals,omitempty" json:"lua_globals,omitempty"` 
	Apparitions_per_level []Apparitions_per_level_Struct  `xml:"apparitions_per_level,omitempty" json:"apparitions_per_level,omitempty"` 
	Orbs_found_thisrun []Orbs_found_thisrun_Struct  `xml:"orbs_found_thisrun,omitempty" json:"orbs_found_thisrun,omitempty"` 
	Flags []Flags_Struct  `xml:"flags,omitempty" json:"flags,omitempty"` 
	Changed_materials []Changed_materials_Struct  `xml:"changed_materials,omitempty" json:"changed_materials,omitempty"` 
}

type Lua_globals_Struct struct {
	E []E_Struct  `xml:"E,omitempty" json:"E,omitempty"` 
}

type E_Struct struct {
	Key string  `xml:"key,attr,omitempty" json:"key,omitempty"` 
	Value string  `xml:"value,attr,omitempty" json:"value,omitempty"` 
}

type Apparitions_per_level_Struct struct {
	Primitive_Struct []string  `xml:"primitive,omitempty" json:"primitive,omitempty"` 
}

type Orbs_found_thisrun_Struct struct {
	Primitive_Struct []string  `xml:"primitive,omitempty" json:"primitive,omitempty"` 
}

type Flags_Struct struct {
	String_Struct []string  `xml:"string,omitempty" json:"string,omitempty"` 
}

type Changed_materials_Struct struct {
	String_Struct []string  `xml:"string,omitempty" json:"string,omitempty"` 
}

type PlayerStatsComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Lives string  `xml:"lives,attr,omitempty" json:"lives,omitempty"` 
	Max_hp string  `xml:"max_hp,attr,omitempty" json:"max_hp,omitempty"` 
	Speed string  `xml:"speed,attr,omitempty" json:"speed,omitempty"` 
}

type TeleportComponent_Struct struct {
	Enabled string  `xml:"_enabled,attr,omitempty" json:"_enabled,omitempty"` 
	Target_x_is_absolute_position string  `xml:"target_x_is_absolute_position,attr,omitempty" json:"target_x_is_absolute_position,omitempty"` 
	Target_x string  `xml:"target.x,attr,omitempty" json:"target.x,omitempty"` 
	Target_y string  `xml:"target.y,attr,omitempty" json:"target.y,omitempty"` 
}

type LifetimeComponent_Struct struct {
	Lifetime string  `xml:"lifetime,attr,omitempty" json:"lifetime,omitempty"` 
}

type LastDailyRunPlayed_Struct struct {
	Seed string  `xml:"seed,attr,omitempty" json:"seed,omitempty"` 
}

type Stats_Struct struct {
	Deaths string  `xml:"deaths,attr,omitempty" json:"deaths,omitempty"` 
	Kills string  `xml:"kills,attr,omitempty" json:"kills,omitempty"` 
	Player_kills string  `xml:"player_kills,attr,omitempty" json:"player_kills,omitempty"` 
	Player_projectile_count string  `xml:"player_projectile_count,attr,omitempty" json:"player_projectile_count,omitempty"` 
	Death_map []Death_map_Struct  `xml:"death_map,omitempty" json:"death_map,omitempty"` 
	Kill_map []Kill_map_Struct  `xml:"kill_map,omitempty" json:"kill_map,omitempty"` 
	BUILD_NAME string  `xml:"BUILD_NAME,attr,omitempty" json:"BUILD_NAME,omitempty"` 
	Stats []Stats_Struct  `xml:"stats,omitempty" json:"stats,omitempty"` 
	Biomes_visited_with_wands string  `xml:"biomes_visited_with_wands,attr,omitempty" json:"biomes_visited_with_wands,omitempty"` 
	Damage_taken string  `xml:"damage_taken,attr,omitempty" json:"damage_taken,omitempty"` 
	Dead string  `xml:"dead,attr,omitempty" json:"dead,omitempty"` 
	Death_count string  `xml:"death_count,attr,omitempty" json:"death_count,omitempty"` 
	Death_pos_x string  `xml:"death_pos.x,attr,omitempty" json:"death_pos.x,omitempty"` 
	Death_pos_y string  `xml:"death_pos.y,attr,omitempty" json:"death_pos.y,omitempty"` 
	Enemies_killed string  `xml:"enemies_killed,attr,omitempty" json:"enemies_killed,omitempty"` 
	Killed_by string  `xml:"killed_by,attr,omitempty" json:"killed_by,omitempty"` 
	Gold string  `xml:"gold,attr,omitempty" json:"gold,omitempty"` 
	Gold_all string  `xml:"gold_all,attr,omitempty" json:"gold_all,omitempty"` 
	Gold_infinite string  `xml:"gold_infinite,attr,omitempty" json:"gold_infinite,omitempty"` 
	Healed string  `xml:"healed,attr,omitempty" json:"healed,omitempty"` 
	Heart_containers string  `xml:"heart_containers,attr,omitempty" json:"heart_containers,omitempty"` 
	Hp string  `xml:"hp,attr,omitempty" json:"hp,omitempty"` 
	Items string  `xml:"items,attr,omitempty" json:"items,omitempty"` 
	Kicks string  `xml:"kicks,attr,omitempty" json:"kicks,omitempty"` 
	Places_visited string  `xml:"places_visited,attr,omitempty" json:"places_visited,omitempty"` 
	Playtime string  `xml:"playtime,attr,omitempty" json:"playtime,omitempty"` 
	Playtime_str string  `xml:"playtime_str,attr,omitempty" json:"playtime_str,omitempty"` 
	Projectiles_shot string  `xml:"projectiles_shot,attr,omitempty" json:"projectiles_shot,omitempty"` 
	Streaks string  `xml:"streaks,attr,omitempty" json:"streaks,omitempty"` 
	Teleports string  `xml:"teleports,attr,omitempty" json:"teleports,omitempty"` 
	Wands_edited string  `xml:"wands_edited,attr,omitempty" json:"wands_edited,omitempty"` 
	World_seed string  `xml:"world_seed,attr,omitempty" json:"world_seed,omitempty"` 
	Biome_baseline []Biome_baseline_Struct  `xml:"biome_baseline,omitempty" json:"biome_baseline,omitempty"` 
	Biomes_visited []Biomes_visited_Struct  `xml:"biomes_visited,omitempty" json:"biomes_visited,omitempty"` 
}

type Death_map_Struct struct {
	E []E_Struct  `xml:"E,omitempty" json:"E,omitempty"` 
}

type Kill_map_Struct struct {
	E []E_Struct  `xml:"E,omitempty" json:"E,omitempty"` 
}

type Biome_baseline_Struct struct {
	Biomes_visited_with_wands string  `xml:"biomes_visited_with_wands,attr,omitempty" json:"biomes_visited_with_wands,omitempty"` 
	Damage_taken string  `xml:"damage_taken,attr,omitempty" json:"damage_taken,omitempty"` 
	Dead string  `xml:"dead,attr,omitempty" json:"dead,omitempty"` 
	Death_count string  `xml:"death_count,attr,omitempty" json:"death_count,omitempty"` 
	Death_pos_x string  `xml:"death_pos.x,attr,omitempty" json:"death_pos.x,omitempty"` 
	Death_pos_y string  `xml:"death_pos.y,attr,omitempty" json:"death_pos.y,omitempty"` 
	Enemies_killed string  `xml:"enemies_killed,attr,omitempty" json:"enemies_killed,omitempty"` 
	Gold string  `xml:"gold,attr,omitempty" json:"gold,omitempty"` 
	Gold_all string  `xml:"gold_all,attr,omitempty" json:"gold_all,omitempty"` 
	Gold_infinite string  `xml:"gold_infinite,attr,omitempty" json:"gold_infinite,omitempty"` 
	Healed string  `xml:"healed,attr,omitempty" json:"healed,omitempty"` 
	Heart_containers string  `xml:"heart_containers,attr,omitempty" json:"heart_containers,omitempty"` 
	Hp string  `xml:"hp,attr,omitempty" json:"hp,omitempty"` 
	Items string  `xml:"items,attr,omitempty" json:"items,omitempty"` 
	Kicks string  `xml:"kicks,attr,omitempty" json:"kicks,omitempty"` 
	Places_visited string  `xml:"places_visited,attr,omitempty" json:"places_visited,omitempty"` 
	Playtime string  `xml:"playtime,attr,omitempty" json:"playtime,omitempty"` 
	Playtime_str string  `xml:"playtime_str,attr,omitempty" json:"playtime_str,omitempty"` 
	Projectiles_shot string  `xml:"projectiles_shot,attr,omitempty" json:"projectiles_shot,omitempty"` 
	Streaks string  `xml:"streaks,attr,omitempty" json:"streaks,omitempty"` 
	Teleports string  `xml:"teleports,attr,omitempty" json:"teleports,omitempty"` 
	Wands_edited string  `xml:"wands_edited,attr,omitempty" json:"wands_edited,omitempty"` 
	World_seed string  `xml:"world_seed,attr,omitempty" json:"world_seed,omitempty"` 
}

type Biomes_visited_Struct struct {
	E []E_Struct  `xml:"E,omitempty" json:"E,omitempty"` 
}