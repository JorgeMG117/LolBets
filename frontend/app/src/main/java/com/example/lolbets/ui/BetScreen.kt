package com.example.lolbets.ui

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.Button
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
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
import androidx.compose.ui.text.input.KeyboardType
import com.example.lolbets.data.BetUiState

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun BetScreen(betState: BetUiState, modifier: Modifier = Modifier) {
    Column (
        verticalArrangement = Arrangement.SpaceEvenly,
        horizontalAlignment = Alignment.CenterHorizontally,
        modifier = Modifier.fillMaxHeight()
    ) {
        //Row (verticalAlignment = Alignment.CenterVertically,
        //horizontalArrangement = Arrangement.SpaceBetween,)

        var button1BackgroundColor by remember { mutableStateOf(0xffffffff) }
        var button2BackgroundColor by remember { mutableStateOf(0xffffffff) }

        //TODO: Añadir mas padding en esta pantalla
        GameCard(
            betState.game,
            button1BackgroundColor,
            {
                if (button1BackgroundColor == 0xff78bbff) {
                    button1BackgroundColor = 0xffffffff
                    button2BackgroundColor = 0xffffffff
                }
                else {
                    button1BackgroundColor = 0xff78bbff
                    button2BackgroundColor = 0xffffffff
                }
                //betState.teamChoice = 1
            },
            button2BackgroundColor,
            {
                if (button2BackgroundColor == 0xff78bbff) {
                    button1BackgroundColor = 0xffffffff
                    button2BackgroundColor = 0xffffffff
                }
                else {
                    button2BackgroundColor = 0xff78bbff
                    button1BackgroundColor = 0xffffffff
                }
            },
            modifier
        )

        var text by remember { mutableStateOf("") }

        OutlinedTextField(
            value = text,
            onValueChange = { text = it },
            label = { Text("Place your bet") },
            keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number)
        )

        Button(onClick = {  }) {
            Text(text = "Place my bet")
        }
    }
}

/*@Preview(showBackground = true)
@Composable
fun BetPreview() {
    BetScreen(BetUiState(Game(Team(R.string.team_name_astralis, R.drawable.astralis), Team(R.string.team_name_fnatic, R.drawable.fnatic), League(R.string.league_name_lec, R.drawable.lec), "10 de junio", 100, 100),0 ))
}*/