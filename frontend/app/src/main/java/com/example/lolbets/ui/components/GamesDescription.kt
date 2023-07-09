package com.example.lolbets.ui.components

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import com.example.lolbets.model.Game

@Composable
fun GameCard(game : Game, modifier: Modifier = Modifier){
    Column {
        Row (
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.SpaceBetween,
            modifier = Modifier
                .fillMaxWidth()
                .background(Color(0xffb0bfd9))
        ){
            Row (
                verticalAlignment = Alignment.CenterVertically,
            ) {
                Text(
                    text = LocalContext.current.getString(game.league.stringResourceId),
                    fontWeight = FontWeight.Bold
                )
                Image(
                    painter = painterResource(game.league.imageResourceId),
                    contentDescription = stringResource(game.league.stringResourceId),
                    modifier = Modifier,
                    //.fillMaxWidth()
                    //.height(194.dp),
                    contentScale = ContentScale.Crop
                )
                Text(text = game.date, fontWeight = FontWeight.Bold)
            }
            Text(
                text = "Best of 1",
                //modifier = modifier.align(alignment = Alignment.End)
            )

        }

        Row(
            horizontalArrangement = Arrangement.SpaceEvenly,
            verticalAlignment = Alignment.CenterVertically,
            modifier = modifier
                .background(Color(0xffb0bfd9))
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
                        text = LocalContext.current.getString(game.team1.stringResourceId)
                            .replace(" ", "\n"),
                        modifier = Modifier.padding(16.dp),
                        style = MaterialTheme.typography.headlineSmall
                    )
                    Image(
                        painter = painterResource(game.team1.imageResourceId),
                        contentDescription = stringResource(game.team1.stringResourceId),
                        modifier = Modifier,
                        //.fillMaxWidth()
                        //.height(194.dp),
                        contentScale = ContentScale.Crop
                    )

                }
                Text(text = "10V-4D", fontWeight = FontWeight.Bold)

                Text(
                    text = "1.2",
                    fontWeight = FontWeight.Bold,
                    modifier = Modifier
                        //.background(Color(0xff78bbff))
                        //.border(BorderStroke(1.dp, Color.Black))
                        .padding(
                            horizontal = 20.dp,
                            vertical = 15.dp
                        )
                )
            }


            Text(text = "VS", fontWeight = FontWeight.Bold)


            Column(
                //verticalArrangement = Arrangement.SpaceEvenly,
                horizontalAlignment = Alignment.CenterHorizontally
            ) {
                Row(
                    verticalAlignment = Alignment.CenterVertically
                ) {
                    Image(
                        painter = painterResource(game.team2.imageResourceId),
                        contentDescription = stringResource(game.team2.stringResourceId),
                        modifier = Modifier,
                        //.fillMaxWidth()
                        //.height(194.dp),
                        contentScale = ContentScale.Crop
                    )
                    Text(
                        text = LocalContext.current.getString(game.team2.stringResourceId)
                            .replace(" ", "\n"),
                        modifier = Modifier.padding(16.dp),
                        style = MaterialTheme.typography.headlineSmall
                    )
                }
                Text(text = "10V-4D", fontWeight = FontWeight.Bold)

                Text(
                    text = "1.8",
                    fontWeight = FontWeight.Bold,
                    modifier = Modifier
                        //.background(Color(0xff78bbff))
                        //.border(BorderStroke(1.dp, Color.Black))
                        .padding(
                            horizontal = 20.dp,
                            vertical = 15.dp
                        )
                )
            }


        }

    }
}

@Composable
fun GamesList(gamesList: List<Game>, contentPadding: PaddingValues, modifier: Modifier = Modifier) {
    LazyColumn(modifier = modifier, contentPadding = contentPadding) {
        items(gamesList) { game ->
            GameCard(
                game = game,
                modifier = Modifier.padding(bottom = 30.dp)
            )
        }
    }
}