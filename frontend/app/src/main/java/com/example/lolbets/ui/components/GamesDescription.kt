package com.example.lolbets.ui.components

import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.MaterialTheme
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
import com.example.lolbets.model.Game
import com.example.lolbets.model.League
import com.example.lolbets.model.Team
import com.example.lolbets.ui.BetScreen
import coil.compose.AsyncImage

@Composable
fun MatchDescription(game : Game, bet1ButtonColor: Long, onBet1Clicked: () -> Unit, bet2ButtonColor: Long, onBet2Clicked: () -> Unit, modifier: Modifier = Modifier){
    Row(
        horizontalArrangement = Arrangement.SpaceEvenly,
        verticalAlignment = Alignment.CenterVertically,
        modifier = modifier
            //.background(Color(0xffb0bfd9))
            .fillMaxWidth()
    ) {

        Column(
            //verticalArrangement = Arrangement.SpaceEvenly,
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            Row(
                verticalAlignment = Alignment.CenterVertically
            ) {
                Text(
                    text = game.team1.code
                        .replace(" ", "\n"),
                    //modifier = Modifier.padding(16.dp),
                    style = MaterialTheme.typography.headlineSmall,
                    fontSize = 15.sp
                )
                /*Image(
                    painter = painterResource(game.team1.imageResourceId),
                    contentDescription = stringResource(game.team1.stringResourceId),
                    //modifier = Modifier.background(Color.Black),
                    //.fillMaxWidth()
                    //.height(194.dp),
                    contentScale = ContentScale.Crop
                )*/
                AsyncImage(
                    model = game.team1.image,
                    contentDescription = game.team1.name,
                    contentScale = ContentScale.Crop,
                    modifier = Modifier
                        .size(50.dp),
                        //.clip(CircleShape)
                )

            }
            Text(text = "10V-4D", fontWeight = FontWeight.Bold)

            Button(
                modifier = Modifier
                    .padding(
                        horizontal = 20.dp,
                        vertical = 15.dp
                    ),
                shape = RoundedCornerShape(10.dp),
                colors = ButtonDefaults.buttonColors(
                    containerColor = Color(bet1ButtonColor)
                ),
                elevation = ButtonDefaults.buttonElevation(
                    defaultElevation = 10.dp,
                    pressedElevation = 15.dp,
                    disabledElevation = 0.dp
                ),
                onClick = { onBet1Clicked() }
            ) {
                val value: Double = 1.0 + game.betsTeam1.toDouble() / game.betsTeam2.toDouble()
                val stringValue = String.format("%.1f", value)
                Text(
                    text = stringValue,
                    fontWeight = FontWeight.Bold,
                    color = Color.Black
                )
            }
        }


        Text(text = "VS", fontWeight = FontWeight.Bold)


        Column(
            //verticalArrangement = Arrangement.SpaceEvenly,
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            Row(
                verticalAlignment = Alignment.CenterVertically
            ) {
                AsyncImage(
                    model = game.team2.image,
                    contentDescription = game.team2.name,
                    contentScale = ContentScale.Crop,
                    modifier = Modifier
                        .size(50.dp),
                )
                Text(
                    text = game.team2.code
                        .replace(" ", "\n"),
                    //modifier = Modifier.padding(16.dp),
                    style = MaterialTheme.typography.headlineSmall,
                    fontSize = 15.sp
                )
            }
            Text(text = "10V-4D", fontWeight = FontWeight.Bold)

            Button(
                modifier = Modifier
                    .padding(
                        horizontal = 20.dp,
                        vertical = 15.dp
                    ),
                shape = RoundedCornerShape(10.dp),
                colors = ButtonDefaults.buttonColors(
                    containerColor = Color(bet2ButtonColor)
                ),
                elevation = ButtonDefaults.buttonElevation(
                    defaultElevation = 10.dp,
                    pressedElevation = 15.dp,
                    disabledElevation = 0.dp
                ),
                onClick = { onBet2Clicked() }
            ) {
                val value: Double = 1.0 + game.betsTeam2.toDouble() / game.betsTeam1.toDouble()
                val stringValue = String.format("%.1f", value)
                Text(
                    text = stringValue,
                    fontWeight = FontWeight.Bold,
                    color = Color.Black
                )
            }
        }


    }
}

@Composable
fun GameCard(game : Game, bet1ButtonColor: Long, onBet1Clicked: () -> Unit, bet2ButtonColor: Long, onBet2Clicked: () -> Unit, modifier: Modifier = Modifier){
    Column(modifier = modifier) {
        Row (
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.SpaceBetween,
            modifier = Modifier
                .fillMaxWidth()
                //.background(Color(0xffb0bfd9))
        ){
            Row (
                verticalAlignment = Alignment.CenterVertically,
            ) {
                AsyncImage(
                    model = game.league.image,
                    contentDescription = game.league.name,
                    contentScale = ContentScale.Crop,
                    modifier = Modifier
                        .size(50.dp)
                        .background(Color.Black)
                        //.clip(CircleShape)
                )
                Text(
                    text = game.league.name,
                    fontWeight = FontWeight.Bold,
                    modifier = Modifier.padding(start = 5.dp)
                )

                // Time of the game
                Text(
                    text = game.date,
                    fontWeight = FontWeight.Bold,
                    modifier = Modifier.padding(start = 16.dp)
                )
            }
            Text(
                text = game.blockName,
                modifier = Modifier.padding(end = 5.dp)
                //modifier = modifier.align(alignment = Alignment.End)
            )

        }

        MatchDescription(game, bet1ButtonColor, onBet1Clicked, bet2ButtonColor, onBet2Clicked, Modifier)

    }
}

@Composable
fun GamesList(gamesList: List<Game>, contentPadding: PaddingValues, onGameClicked: (Game) -> Unit, modifier: Modifier = Modifier) {
    LazyColumn(modifier = modifier, contentPadding = contentPadding) {
        items(gamesList) { game ->
            GameCard(
                game = game,
                bet1ButtonColor = 0xffffffff,
                onBet1Clicked = {},
                bet2ButtonColor = 0xffffffff,
                onBet2Clicked = {},
                modifier = Modifier
                    .padding(bottom = 30.dp)
                    .clickable { onGameClicked(game) }
            )
        }
    }
}

/*@Preview(showBackground = true)
@Composable
fun GameDescription() {
    GameCard(Game(
        Team("Fnatic", "Fnatic", "https://am-a.akamaihd.net/image?resize=140:&f=http%3A%2F%2Fstatic.lolesports.com%2Fteams%2F1631819669150_fnc-2021-worlds.png"),
        Team("Fnatic", "Fnatic", "https://am-a.akamaihd.net/image?resize=140:&f=http%3A%2F%2Fstatic.lolesports.com%2Fteams%2F1631819669150_fnc-2021-worlds.png"),
        League("LEC", "LEC", "EMEA", "https://am-a.akamaihd.net/image?resize=120:&f=http%3A%2F%2Fstatic.lolesports.com%2Fleagues%2F1592516184297_LEC-01-FullonDark.png"),
        "10 de junio", 100, 0, 0, "Semifinal", "Best of 1"), 0xffffffff, {}, 0xffffffff, {})

}*/