package com.example.agemobile;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.text.Editable;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;
import android.widget.Toast;

import com.google.android.material.textfield.TextInputEditText;

import agemobile.Agemobile;

public class MainActivity extends AppCompatActivity {
    age.X25519Identity identity;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        TextInputEditText encryptInput = findViewById(R.id.textToEncrypt);

        TextView keyView = findViewById(R.id.keysView);
        TextView encryptResults = findViewById(R.id.encryptResults);
        TextView decryptResults = findViewById(R.id.decryptText);

        Button generateKey = findViewById(R.id.generateKey);
        Button encryptBtn = findViewById(R.id.encryptBtn);
        Button decryptBtn = findViewById(R.id.decryptBtn);

        generateKey.setOnClickListener(view -> {
            try {
                identity = Agemobile.generateX25519Identity();
                String txt = "Private Key:\n"+identity.string()+"\nPublic Key:\n"+identity.recipient().string();
                keyView.setText(txt);
            } catch (Exception e) {
                e.printStackTrace();
            }

        });

        encryptBtn.setOnClickListener(view -> {
            Editable editable = encryptInput.getText();
            if(editable != null && editable.length() > 0) {
                if(identity == null) {
                    Toast.makeText(view.getContext(), "Please generate key first",Toast.LENGTH_LONG).show();
                    return;
                }
                String text = editable.toString();
                try {
                    // public key, text and with armor
                    String results = Agemobile.encrypt(identity.recipient().string(), text, true);
                    encryptResults.setText(results);
                } catch (Exception e) {
                    e.printStackTrace();
                }
            }else {
                Toast.makeText(view.getContext(), "Please input some text to encrypt",Toast.LENGTH_LONG).show();
            }
        });

        decryptBtn.setOnClickListener(view -> {
            String txt = encryptResults.getText().toString();
            if(!txt.isEmpty()) {
                if(identity == null) {
                    Toast.makeText(view.getContext(), "Please generate key first",Toast.LENGTH_LONG).show();
                    return;
                }

                try {
                    // private key and encrypted text
                    String results = Agemobile.decrypt(identity.string(), txt);
                    decryptResults.setText(results);
                } catch (Exception e) {
                    e.printStackTrace();
                }

            }else {
                Toast.makeText(view.getContext(), "Please encrypt some text",Toast.LENGTH_LONG).show();
            }

        });



    }
}