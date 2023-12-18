package com.example.lolbets.ui

import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.Image
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Card
import androidx.compose.material3.CardDefaults
import androidx.compose.material3.ElevatedCard
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import com.example.lolbets.R
import com.example.lolbets.model.ActiveBets
import com.example.lolbets.model.User
import com.example.lolbets.ui.components.ErrorScreen
import com.example.lolbets.ui.components.GamesList
import com.example.lolbets.ui.components.LoadingScreen
import com.example.lolbets.viewmodel.ActiveBetsUiState
import com.example.lolbets.viewmodel.GameUiState

@Composable
fun BetsSummaryScreen(userBets:  List<ActiveBets>, contentPadding: PaddingValues, modifier: Modifier = Modifier) {
    //Divide the user bets into, active and recent
    val activeBets: List<ActiveBets>
    val recentBets: List<ActiveBets>

    if(userBets.isEmpty()) {
        activeBets = userBets
        recentBets = userBets
    } else {
        val index = userBets.indexOfFirst { it.completed > 0 }

        if(index == -1) {//All the elements are just active bets
            activeBets = userBets
            recentBets = emptyList()
        }
        else {
            activeBets = userBets.take(index)
            recentBets = userBets.drop(index)
        }
    }
    val combinedList = mutableListOf<Any>()
    combinedList.add("Active Bets")
    combinedList.addAll(activeBets)
    combinedList.add("Recent Results")
    combinedList.addAll(recentBets)

    LazyColumn(
        horizontalAlignment = Alignment.CenterHorizontally,
        modifier = Modifier.fillMaxWidth()
    ) {
        items(combinedList) { item ->
            when (item) {
                is String -> {
                    // This is a header
                    Text(
                        text = item,
                        fontWeight = FontWeight.Bold,
                        fontSize = 30.sp
                    )
                }

                is ActiveBets -> {
                    // This is a bet item
                    val bet = item
                    var won = false
                    if(bet.bet.team) {
                        won = bet.completed == 2
                    }
                    else {
                        won = bet.completed == 1
                    }

                    Card(
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(16.dp),
                        elevation = CardDefaults.cardElevation(
                            defaultElevation = 6.dp
                        )
                    ) {
                        Column(
                            modifier = Modifier
                                .padding(16.dp)
                        ) {
                            Row(
                                modifier = Modifier
                                    .fillMaxWidth()
                                    .padding(bottom = 8.dp),
                                horizontalArrangement = Arrangement.SpaceBetween
                            ) {
                                Text(
                                    text = "${bet.team1} vs ${bet.team2}",
                                    fontWeight = FontWeight.Bold,
                                    fontSize = 18.sp
                                )
                                if (bet.completed != 0) {
                                    Text(
                                        text = "Completed",
                                        color = if (won) Color.Green else Color.Red,
                                        fontWeight = FontWeight.Bold,
                                        fontSize = 16.sp
                                    )
                                }
                            }
                            Spacer(modifier = Modifier.height(8.dp))
                            Text(text = bet.league, fontSize = 16.sp)
                            Spacer(modifier = Modifier.height(4.dp))
                            Text(text = "Bet Value: ${bet.bet.value}", fontSize = 16.sp)
                            Spacer(modifier = Modifier.height(4.dp))
                            Text(text = "Odds: ${bet.bet.odds}", fontSize = 16.sp)
                        }
                    }
                }
            }
        }
    }
        //Esto tiene que salir de la BD, ya que es en la BD donde se guarda quien ha ganado
        //Una vez que acabe el game, se meteran las activeBets en la BD
        //Aun asi aun no se sabe el resultado del partido, ya que guardamos el game en la BD
        //pero el resultado aun no se sabe, hay que esperar a que se ejecute una funcion en el
        //backend que pregunte a la api y actualize la BD
}

@Composable
fun BetsSummary(activeBetsUiState: ActiveBetsUiState, contentPadding: PaddingValues, modifier: Modifier = Modifier) {
    when (activeBetsUiState) {
        is ActiveBetsUiState.Loading -> LoadingScreen( modifier = modifier.fillMaxSize())

        is ActiveBetsUiState.Success -> BetsSummaryScreen(
            userBets = activeBetsUiState.activeBets,
            contentPadding = contentPadding
        )

        is ActiveBetsUiState.Error -> ErrorScreen( modifier = modifier.fillMaxSize())
    }


}

