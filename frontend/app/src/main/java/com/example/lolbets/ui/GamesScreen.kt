package com.example.lolbets.ui

import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.lolbets.data.GamesData
import com.example.lolbets.model.Game
import com.example.lolbets.ui.components.GamesList
import com.example.lolbets.viewmodel.GameUiState


@Composable
fun HomeScreen(
    gameUiState: GameUiState, contentPadding: PaddingValues, onGameClicked: (Game) -> Unit, modifier: Modifier = Modifier
) {
    when (gameUiState) {
        is GameUiState.Loading -> LoadingScreen( modifier = modifier.fillMaxSize())

        is GameUiState.Success -> GamesList(
            gamesList = gameUiState.games,
            contentPadding = contentPadding,
            onGameClicked = onGameClicked,
        )

        is GameUiState.Error -> ErrorScreen( modifier = modifier.fillMaxSize())
    }
}

/**
 * ResultScreen displaying number of photos retrieved.
 */
@Composable
fun ResultScreen(games: List<Game>, modifier: Modifier = Modifier) {
    LazyColumn(
        modifier = Modifier.fillMaxSize()
    ) {
        items(games) { game ->
            Text(text = game.date)
        }
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