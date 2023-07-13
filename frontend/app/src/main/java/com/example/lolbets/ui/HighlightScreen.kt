package com.example.lolbets.ui

import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.lolbets.data.GamesData
import com.example.lolbets.ui.components.GamesList


@Composable
fun HighlightScreen(contentPadding: PaddingValues, modifier: Modifier = Modifier) {
    /*GamesList(
        gamesList = GamesData().loadGames(),
        contentPadding = contentPadding
    )*/
    Text(text = "HighlightScreen")
}
@Preview(showBackground = true)
@Composable
fun HighlightPreview() {
    HighlightScreen(contentPadding = PaddingValues(10.dp))
}