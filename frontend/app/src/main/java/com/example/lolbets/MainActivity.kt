package com.example.lolbets

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn

import androidx.compose.foundation.lazy.items

import androidx.compose.material3.Card
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.lolbets.data.LeagueData
import com.example.lolbets.model.League
import com.example.lolbets.ui.theme.LolBetsTheme
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.AccountCircle
import androidx.compose.material.icons.filled.ArrowBack
import androidx.compose.material.icons.filled.Favorite
import androidx.compose.material.icons.filled.KeyboardArrowLeft
import androidx.compose.material.icons.filled.Menu
import androidx.compose.material.icons.filled.Person
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.Scaffold
import androidx.compose.material3.TopAppBar
import androidx.compose.ui.Alignment
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextOverflow
import com.example.lolbets.data.GamesData
import com.example.lolbets.model.Game
import com.example.lolbets.model.Team


class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            LolBetsTheme {
                // A surface container using the 'background' color from the theme
                Surface(
                    modifier = Modifier.fillMaxSize(),
                    color = MaterialTheme.colorScheme.background
                ) {
                    LeaguesScreen()
                }
            }
        }
    }
}

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
fun LeagueCard(league: League, modifier: Modifier = Modifier) {
    Card(modifier = modifier) {
        Row(
            verticalAlignment = Alignment.CenterVertically,
            modifier = modifier
                .background(Color(0xffb0bfd9))
                .fillMaxWidth()
        ) {
            Image(
                painter = painterResource(league.imageResourceId),
                contentDescription = stringResource(league.stringResourceId),
                modifier = Modifier,
                    //.fillMaxWidth()
                    //.height(194.dp),
                contentScale = ContentScale.Crop
            )
            Text(
                text = LocalContext.current.getString(league.stringResourceId),
                modifier = Modifier.padding(16.dp),
                style = MaterialTheme.typography.headlineSmall
            )
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
internal fun LolBetsTopAppBar(modifier: Modifier = Modifier) {
    TopAppBar(
        modifier = modifier,
        title = {
            Text(
                stringResource(R.string.app_name),
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        },
        navigationIcon = {
            IconButton(onClick = { /* doSomething() */ }) {
                Icon(
                    imageVector = Icons.Filled.KeyboardArrowLeft,
                    contentDescription = "Localized description"
                )
            }
        },
        actions = {
            Text(text = "10$")
            IconButton(onClick = { /* doSomething() */ }) {
                Icon(
                    imageVector = Icons.Filled.AccountCircle,
                    contentDescription = "Localized description"
                )
            }
        }
    )

}

@Composable
internal fun LolBetsBottomAppBar(modifier: Modifier = Modifier) {
    NavigationBar {

    }
}


@Composable
fun LeaguesList(leaguesList: List<League>, modifier: Modifier = Modifier) {
    LazyColumn(modifier = modifier) {
        items(leaguesList) { league ->
            LeagueCard(
                league = league,
                modifier = Modifier.padding(8.dp)
            )
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

@Composable
fun UserCard(league: League, modifier: Modifier = Modifier) {
    Card(modifier = modifier) {
        Row(
            verticalAlignment = Alignment.CenterVertically,
            modifier = modifier
                .background(Color(0xffb0bfd9))
                .fillMaxWidth()
        ) {
            Image(
                painter = painterResource(league.imageResourceId),
                contentDescription = stringResource(league.stringResourceId),
                modifier = Modifier,
                //.fillMaxWidth()
                //.height(194.dp),
                contentScale = ContentScale.Crop
            )
            Text(
                text = LocalContext.current.getString(league.stringResourceId),
                modifier = Modifier.padding(16.dp),
                style = MaterialTheme.typography.headlineSmall
            )
        }
    }
}

@Composable
fun StandingScreen(usersList: List<Game>, contentPadding: PaddingValues, modifier: Modifier = Modifier) {
    LazyColumn(modifier = modifier, contentPadding = contentPadding) {
        items(usersList) { user ->
            GameCard(
                game = user,
                modifier = Modifier.padding(bottom = 30.dp)
            )
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun LeaguesScreen(modifier: Modifier = Modifier) {
    Scaffold(
        topBar = {
            LolBetsTopAppBar(modifier)
        },
        bottomBar = {
            LolBetsBottomAppBar(modifier)
        }
    ) { it ->
        GamesList(
            gamesList = GamesData().loadGames(),
            contentPadding = it
        )
    }

    /*
    LeaguesList(
        leaguesList = LeagueData().loadLeagues()

        GamesList(
            gamesList = GamesData().loadGames(),
            contentPadding = it
        )
    )*/


}

@Preview(showBackground = true)
@Composable
fun GreetingPreview() {
    LolBetsTheme {
        LeaguesScreen()
    }
}