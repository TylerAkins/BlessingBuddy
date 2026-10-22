local _, addon = ...

addon.CanUseSalvation = function(isSelf, isPlayer, inGroup)
	return isSelf or (isPlayer and inGroup)
end
