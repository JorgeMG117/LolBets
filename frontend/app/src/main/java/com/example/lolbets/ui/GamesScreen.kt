package com.example.lolbets.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.lolbets.data.GamesData
import com.example.lolbets.model.Game
import com.example.lolbets.ui.components.GamesList


@Composable
fun HomeScreen(
    gameUiState: GameUiState, modifier: Modifier = Modifier
) {
    when (gameUiState) {
        is GameUiState.Loading -> LoadingScreen(modifier = modifier.fillMaxSize())
        is GameUiState.Success -> ResultScreen(
            gameUiState.games, modifier = modifier.fillMaxWidth()
        )

        is GameUiState.Error -> ErrorScreen( modifier = modifier.fillMaxSize())
    }
}

/**
 * ResultScreen displaying number of photos retrieved.
 */
@Composable
fun ResultScreen(games: String, modifier: Modifier = Modifier) {
    Box(
        contentAlignment = Alignment.Center,
        modifier = modifier
    ) {
        Text(text = games)
    }
}

@Composable
fun LoadingScreen(modifier: Modifier = Modifier) {
    Text(text = "Loading")
}

/**
 * The home screen displaying error message with re-attempt button.
 */
@Composable
fun ErrorScreen(modifier: Modifier = Modifier) {
    Text(text = "Error")
}



@Composable
fun GamesScreen(contentPadding: PaddingValues, onGameClicked: (Game) -> Unit, modifier: Modifier = Modifier) {
    GamesList(
        gamesList = GamesData().loadGames(),
        contentPadding = contentPadding,
        onGameClicked = onGameClicked,
    )
}
@Preview(showBackground = true)
@Composable
fun GamesPreview() {
    GamesScreen(
        contentPadding = PaddingValues(0.dp),
        onGameClicked = {}
    )
}