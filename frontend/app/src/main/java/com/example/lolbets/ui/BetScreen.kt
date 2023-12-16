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
import com.example.lolbets.model.ActiveBets
import com.example.lolbets.model.Bet
import com.example.lolbets.ui.components.getOddsTeam1
import com.example.lolbets.ui.components.getOddsTeam2


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

        //TODO Quiza añadir que se elimine el mensaje de bet succesful
        //El mensaje de bet successful deveria desaparecer cuando cambias de pantalla(se cierra socket) o
        // se escribe un nuevo valor, selecciona nuevo equipo
        OutlinedTextField(
            value = betValue,
            onValueChange = {
                //betState.TODO Limpiar mensaje de apuesta correcta cuando se escriba algo
                betValue = it },
            label = { Text("Place your bet") },
            keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number)
        )

        Button(
            onClick = {
                var placeBet = true
                var betValueInt = 0
                try {
                    betValueInt = betValue.toInt()
                    if(betValueInt <= 0) {
                        placeBet = false
                    }
                } catch (e: NumberFormatException) {
                    placeBet = false
                }
                var team = true
                if (button1BackgroundColor == 0xff78bbff) {
                    team = false
                }
                else if (button2BackgroundColor == 0xff78bbff) {
                    team = true
                }
                else {//No team selected
                    placeBet = false
                }

                if(placeBet) {
                    val odds : Double
                    if(team) {
                        odds = getOddsTeam2(betState.game)
                    }
                    else {
                        odds = getOddsTeam1(betState.game)
                    }
                    //println(betState.userId)
                    onBetPlaced(Bet(betValueInt, team, betState.userId, betState.game.id, odds))
                }
                else {
                    //TODO Show error text
                }
                betValue = "" //Clean input field
            },
            //enabled = betState.isConnected
        ) {
            Text(text = "Place my bet")
        }
        if(betState.betSucceed) {
            Text(text = "Bet placed successfully")
        }
    }
}

/*@Preview(showBackground = true)
@Composable
fun BetPreview() {
    BetScreen(BetUiState(Game(Team(R.string.team_name_astralis, R.drawable.astralis), Team(R.string.team_name_fnatic, R.drawable.fnatic), League(R.string.league_name_lec, R.drawable.lec), "10 de junio", 100, 100),0 ))
}*/