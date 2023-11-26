package com.example.lolbets.ui

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.Button
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import com.example.lolbets.ui.components.GameCard
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.ui.text.input.KeyboardType
import com.example.lolbets.data.BetUiState
import com.example.lolbets.model.Bet


@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun BetScreen(betState: BetUiState, onBetPlaced: (Bet) -> Unit, modifier: Modifier = Modifier) {
    Column (
        verticalArrangement = Arrangement.SpaceEvenly,
        horizontalAlignment = Alignment.CenterHorizontally,
        modifier = Modifier.fillMaxHeight()
    ) {
        //Row (verticalAlignment = Alignment.CenterVertically,
        //horizontalArrangement = Arrangement.SpaceBetween,)
        //PlaceBet()

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

        var betValue by remember { mutableStateOf("") }
        //var bet by remember { mutableStateOf(Bet(0, false, 0, 0)) }

        OutlinedTextField(
            value = betValue,
            onValueChange = { betValue = it },
            label = { Text("Place your bet") },
            keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number)
        )

        Button(
            onClick = {
                println(betValue)
                var team = true
                if (button1BackgroundColor == 0xff78bbff) {
                    team = false
                }
                onBetPlaced(Bet(betValue.toInt(), team, 0, 10, 1.5))//TODO
            },
            //enabled = betState.isConnected
        ) {
            Text(text = "Place my bet")
        }
    }
}

/*@Preview(showBackground = true)
@Composable
fun BetPreview() {
    BetScreen(BetUiState(Game(Team(R.string.team_name_astralis, R.drawable.astralis), Team(R.string.team_name_fnatic, R.drawable.fnatic), League(R.string.league_name_lec, R.drawable.lec), "10 de junio", 100, 100),0 ))
}*/