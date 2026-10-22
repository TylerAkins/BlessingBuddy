local addon = {}
assert(loadfile("SalvationRules.lua"))("BlessingBuddy", addon)

local cases = {
	{ "self outside a group", true, false, false, true },
	{ "self in a group", true, true, true, true },
	{ "group player", false, true, true, true },
	{ "player outside a group", false, true, false, false },
	{ "non-player group unit", false, false, true, false },
}

for _, case in ipairs(cases) do
	local name, isSelf, isPlayer, inGroup, want = case[1], case[2], case[3], case[4], case[5]
	assert(addon.CanUseSalvation(isSelf, isPlayer, inGroup) == want, name)
end
