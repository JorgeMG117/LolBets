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
import androidx.compose.foundation.shape.CircleShape
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
    Column (
        //verticalArrangement = Arrangement.SpaceEvenly,
        horizontalAlignment = Alignment.CenterHorizontally,
        modifier = modifier.fillMaxWidth()
    ) {
        //Divide the user bets into, active and recent
        val index = userBets.indexOfFirst { it.completed > 0 }

        val activeBets = userBets.take(index)
        val recentBets = userBets.drop(index)

        Text(
            text = "Active Bets",
            fontWeight = FontWeight.Bold,
            fontSize = 30.sp
        )
        LazyColumn {
            items(activeBets) { bet ->
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
        //Tal y como lo tengo, las apuestas del usuario, solo se guardan en la BD cuando
        //se ha acabado el partido
        //Tengo las apuestas en una lista mientras tanto
        //Habra que hacer un get que me devuelva las ag.activeBets de un usuario

        //Que quiero mostrar?
        //Valor, Equipo,
        Text(
            text = "Recent Results",
            fontWeight = FontWeight.Bold,
            fontSize = 30.sp
        )
        LazyColumn {
            items(recentBets) { bet ->
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
                            Text(
                                text = "Completed",
                                color = if (won) Color.Green else Color.Red,
                                fontWeight = FontWeight.Bold,
                                fontSize = 16.sp
                            )
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
        //Esto tiene que salir de la BD, ya que es en la BD donde se guarda quien ha ganado
        //Una vez que acabe el game, se meteran las activeBets en la BD
        //Aun asi aun no se sabe el resultado del partido, ya que guardamos el game en la BD
        //pero el resultado aun no se sabe, hay que esperar a que se ejecute una funcion en el
        //backend que pregunte a la api y actualize la BD
    }
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

