package com.example.lolbets.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.lolbets.data.GamesData
import com.example.lolbets.ui.components.GamesList


@Composable
fun GamesScreen(contentPadding: PaddingValues, modifier: Modifier = Modifier) {
    GamesList(
        gamesList = GamesData().loadGames(),
        contentPadding = contentPadding
    )
}
@Preview(showBackground = true)
@Composable
fun GamesPreview() {
    GamesScreen(contentPadding = PaddingValues(10.dp))
}