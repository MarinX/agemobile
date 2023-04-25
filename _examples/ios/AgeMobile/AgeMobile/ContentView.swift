//
//  ContentView.swift
//  AgeMobile
//
//  Created by Markus Zoppelt on 24.04.23.
//

import SwiftUI
import Age

struct ContentView: View {
    
    @State private var input: String = "my secret"
    @State private var encrypted: String = ""
    @State private var decrypted: String = ""
    @State private var identity = AgemobileGenerateX25519Identity(nil)
    
    private let pasteboard = UIPasteboard.general
    
    var body: some View {
        
        VStack {
            Button("KeyGen", action: {
                identity = AgemobileGenerateX25519Identity(nil)
            })
            
            HStack {
                Button {
                    pasteboard.string = identity?.string()
                } label: {
                    Label("Copy private key", systemImage: "doc.on.doc").foregroundColor(.red)
                }.padding(2)
                Button {
                    pasteboard.string = identity?.recipient()?.string()
                } label: {
                    Label("Copy public key", systemImage: "doc.on.doc").foregroundColor(.green)
                }
            }
            
            TextField("Secret", text: $input, axis: .vertical)
                .textInputAutocapitalization(.never)
                .disableAutocorrection(true)
                .border(.primary)
            
            Button("Encrypt", action: {
                encrypted = AgemobileEncrypt(identity?.recipient()?.string(), input, true, nil)
            })
            
            TextField("Encrypted", text: $encrypted, axis: .vertical)
                .textInputAutocapitalization(.never)
                .disableAutocorrection(true)
                .border(.primary)
            
            Button("Decrypt", action: {
                decrypted = AgemobileDecrypt(identity?.string(), encrypted,nil)
            })
            
            TextField("Decrypted", text: $decrypted, axis: .vertical)
                .textInputAutocapitalization(.never)
                .disableAutocorrection(true)
                .border(.primary)
            
        }}
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
