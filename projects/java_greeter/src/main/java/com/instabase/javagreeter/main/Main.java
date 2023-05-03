package com.instabase.javagreeter.main;

import com.instabase.javagreeter.greeter.Greeter;

class Main {
    public static void main(String[] args) {
        Greeter greeter = new Greeter();
        System.out.println(greeter.greet());
    }
}