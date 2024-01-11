package translations

func initFrenchTranslation() Translation {
	translation := createTranslation()

	translation.put("requires-js", "Ce site a besoin de JavaScript pour fonctionner correctement.")

	translation.put("start-the-game", "Démarrer la partie")
	translation.put("start", "Démarrer")
	translation.put("game-not-started-title", "La partie n'a pas encore commencé")
	translation.put("waiting-for-host-to-start", "Veuillez attendre que l'hôte lance la partie")

	translation.put("round", "Tour")
	translation.put("toggle-soundeffects", "Activer/Désactiver les effets sonores")
	translation.put("change-your-name", "Changer le pseudo")
	translation.put("randomize", "Aléatoire")
	translation.put("apply", "Appliquer")
	translation.put("save", "Sauvegarder")
	translation.put("toggle-fullscreen", "Mettre en plein écran")
	translation.put("show-help", "Voir l'aide")
	translation.put("votekick-a-player", "Voter pour dégager un·e joueureuse")

	translation.put("last-turn", "(Dernier tour: %s)")

	translation.put("drawer-kicked", "Puisque læ joueureuse expulsé·e était en train de déssiner, personne ne prend de points ce tour.")
	translation.put("self-kicked", "Vous avez été expulsé·e")
	translation.put("kick-vote", "(%s/%s) joueureuses ont voté pour expulser %s.")
	translation.put("player-kicked", "Læ joueureuse a été expulsé·e.")
	translation.put("owner-change", "%s est le nouvel hôte de la partie.")

	translation.put("change-lobby-settings-tooltip", "Modifier les paramètres de la partie")
	translation.put("change-lobby-settings-title", "Paramètres de la partie")
	translation.put("lobby-settings-changed", "Paramètres de la partie modifiés")
	translation.put("advanced-settings", "Paramètres avancés")
	translation.put("word-language", "Langage des mots")
	translation.put("drawing-time-setting", "Temps de dessin")
	translation.put("rounds-setting", "Tours")
	translation.put("max-players-setting", "Nombre max de joueureuses")
	translation.put("public-lobby-setting", "Partie publique")
	translation.put("word-groups", "Word Groups")
	translation.put("players-per-ip-limit-setting", "Limite de joueureuses par addresse IP")
	translation.put("save-settings", "Enregistrer les paramètres")
	translation.put("input-contains-invalid-data", "L'entrée contient des données invalides:")
	translation.put("please-fix-invalid-input", "Veuillez corriger votre entrée invalide et réessayez")
	translation.put("no-wordgroup-selected", "Veuillez sélectionner au moins un groupe de mots")
	translation.put("create-lobby", "Créer la partie")

	translation.put("players", "Joueureuses")
	translation.put("refresh", "Rafraichir")
	translation.put("join-lobby", "Rejoindre la partie")

	translation.put("message-input-placeholder", "Écrivez vos propositions et vos messages ici")

	translation.put("choose-a-word", "Choisissez un mot")
	translation.put("waiting-for-word-selection", "En attente de la séléction du mot")
	// This one doesn't use %s, since we want to make one part bold.
	translation.put("is-choosing-word", "choisit un mot.")

	translation.put("close-guess", "'%s' est très proche.")
	translation.put("correct-guess", "Vous avez bien deviné le mot :)")
	translation.put("correct-guess-other-player", "'%s' a bien deviné le mot.")
	translation.put("round-over", "Fin du tour, aucun mot n'a été choisi")
	translation.put("round-over-no-word", "Fin du tour, le mot était '%s'.")
	translation.put("game-over-win", "Félicitation, vous avez gagné !")
	translation.put("game-over-tie", "Égalité !")
	translation.put("game-over", "Vous êtes %s. avec %s points")

	translation.put("change-active-color", "Change votre couleur active")
	translation.put("use-pencil", "Utiliser le crayon")
	translation.put("use-eraser", "Utiliser une gomme")
	translation.put("use-fill-bucket", "Utiliser le seau de remplissage (Remplis la zone avec la couleur séléctionnée)")
	translation.put("change-pencil-size-to", "Mettre la taille du crayon/gomme à %s")
	translation.put("clear-canvas", "Nettoyer l'écran")
	translation.put("undo", "Annuler la dernière modification effectuée (ne fonctionne pas après \""+translation.Get("clear-canvas")+"\")")

	translation.put("connection-lost", "Connection perdue !")
	translation.put("connection-lost-text", "Tentative de reconnection"+
		" ...\n\nAssurez-vous que votre connection fonctionne.\nSi le "+
		"probleme persiste, contactez læ gestionnaire du site.")
	translation.put("error-connecting", "Erreur de connection au serveur")
	translation.put("error-connecting-text",
		"Scribble.rs ne peut pas établir de connection socket.\n\nVotre connection internet "+
			"a pourtant l'air de fonctionner, ou bien le\nserveur ou votre pare-feu n'a "+
			"pas été configuré correctement.\n\nPour réessayez, rafraichissez la page.")

	translation.put("message-too-long", "Votre message est trop long.")

	// Help dialog
	translation.put("controls", "Contrôles")
	translation.put("pencil", "Crayon")
	translation.put("eraser", "Gomme")
	translation.put("fill-bucket", "Seau de remplissage")
	translation.put("switch-tools-intro", "Vous pouvez changer d'outil avec les raccourcis")
	translation.put("switch-pencil-sizes", "Vous pouvez aussi changer la taille des outils avec les touches %s et %s.")

	// Generic words
	// "close" as in "closing the window"
	translation.put("close", "Fermer")
	translation.put("no", "Non")
	translation.put("yes", "Oui")
	translation.put("system", "Systeme")

	translation.put("source-code", "Code Source")
	translation.put("help", "Aide")
	translation.put("contact", "Contact")
	translation.put("submit-feedback", "Feedback")
	translation.put("stats", "Statut")

	RegisterTranslation("fr", translation)
	RegisterTranslation("fr-FR", translation)

	return translation
}
