package com.example.lolbets.ui

import android.app.Activity
import android.content.Intent
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.ActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.core.app.ActivityCompat.startActivityForResult
import com.google.android.gms.auth.api.signin.GoogleSignIn
import com.google.android.gms.auth.api.signin.GoogleSignInAccount
import com.google.android.gms.auth.api.signin.GoogleSignInClient
import com.google.android.gms.tasks.Task


@Composable
fun LoginScreen(clientIdtest: GoogleSignInClient, contentPadding: PaddingValues, modifier: Modifier = Modifier) {
    //Text(text = "Login screen")
    /*val startForResult =
        rememberLauncherForActivityResult(ActivityResultContracts.StartActivityForResult()) { result: ActivityResult ->
            if (result.resultCode == Activity.RESULT_OK) {
                val intent = result.data
                if (result.data != null) {
                    val task: Task<GoogleSignInAccount> =
                        GoogleSignIn.getSignedInAccountFromIntent(intent)
                    //handleSignInResult(task)
                }
            }
        }
    Button(
        onClick = {
            val signInIntent: Intent = clientIdtest.getSignInIntent()
            startActivityForResult(signInIntent, RC_SIGN_IN)
        },
        modifier = Modifier
            .fillMaxWidth()
            .padding(start = 16.dp, end = 16.dp, top = 80.dp),
        shape = RoundedCornerShape(6.dp),
        colors = ButtonDefaults.buttonColors(
            //backgroundColor = Color.Black,
            contentColor = Color.White
        )
    ) {
        /*Image(
            painter = painterResource(id = R.drawable.ic_logo_google),
            contentDescription = ""
        )*/
        Text(text = "Sign in with Google", modifier = Modifier.padding(6.dp))
    }
    */
    /*Button(onClick = { result = (1..6).random() }) {
        Text(text = "Sign in")
    }*/



    /*val startForResult =
            rememberLauncherForActivityResult(ActivityResultContracts.StartActivityForResult()) { result: ActivityResult ->
                if (result.resultCode == Activity.RESULT_OK) {
                    val intent = result.data
                    if (result.data != null) {
                        val task: Task<GoogleSignInAccount> =
                            GoogleSignIn.getSignedInAccountFromIntent(intent)
                        //handleSignInResult(task)
                    }
                }
            }

        Button(
            onClick = {
                startForResult.launch(clientIdtest?.signInIntent)
            },
            modifier = Modifier
                .fillMaxWidth()
                .padding(start = 16.dp, end = 16.dp),
            shape = RoundedCornerShape(6.dp),
        ) {
            Text(text = "Sign in with Google", modifier = Modifier.padding(6.dp))
        }*/
}
@Preview(showBackground = true)
@Composable
fun LoginPreview() {
    //LoginScreen(contentPadding = PaddingValues(10.dp))
}