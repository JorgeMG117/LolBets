package com.example.lolbets.ui

import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.lolbets.data.GamesData
import com.example.lolbets.model.Game
import com.example.lolbets.ui.components.ErrorScreen
import com.example.lolbets.ui.components.GamesList
import com.example.lolbets.ui.components.LoadingScreen
import com.example.lolbets.viewmodel.GameUiState

@Composable
fun HighlightScreen(
    gameUiState: GameUiState, contentPadding: PaddingValues, onGameClicked: (Game) -> Unit, modifier: Modifier = Modifier
) {
    when (gameUiState) {
        is GameUiState.Loading -> LoadingScreen( modifier = modifier.fillMaxSize())

        is GameUiState.Success -> {
            //Get 10 most betted games
            val sortedGames = gameUiState.games.sortedByDescending { it.betsTeam1 + it.betsTeam2 }
            val highlightGames = sortedGames.take(10)

            GamesList(
                gamesList = highlightGames,
                contentPadding = contentPadding,
                onGameClicked = onGameClicked,
            )
        }

        is GameUiState.Error -> ErrorScreen( modifier = modifier.fillMaxSize())
    }
}

/*@Preview(showBackground = true)
@Composable
fun HighlightPreview() {
    HighlightScreen(contentPadding = PaddingValues(10.dp))
}*/