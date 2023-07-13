package com.example.lolbets.ui

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.material3.Button
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.lolbets.R
import com.example.lolbets.data.GamesData
import com.example.lolbets.model.Game
import com.example.lolbets.model.League
import com.example.lolbets.model.Team
import com.example.lolbets.ui.components.GameCard
import com.example.lolbets.ui.components.GamesList
import com.example.lolbets.ui.components.MatchDescription
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.runtime.mutableStateOf

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun BetScreen(contentPadding: PaddingValues, modifier: Modifier = Modifier) {
    val game: Game = Game(Team(R.string.team_name_astralis, R.drawable.astralis), Team(R.string.team_name_fnatic, R.drawable.fnatic), League(R.string.league_name_lec, R.drawable.lec), "10 de junio", 100, 100)

    Column (
        verticalArrangement = Arrangement.SpaceEvenly,
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        //Row (verticalAlignment = Alignment.CenterVertically,
        //horizontalArrangement = Arrangement.SpaceBetween,)

        GameCard(game, modifier)

        var text by remember { mutableStateOf("") }

        OutlinedTextField(
            value = text,
            onValueChange = { text = it },
            label = { Text("Label") }
        )

        Button(onClick = { /*TODO*/ }) {
            Text(text = "Place my bet")
        }
    }
}

@Preview(showBackground = true)
@Composable
fun BetPreview() {
    BetScreen(contentPadding = PaddingValues(10.dp))
}