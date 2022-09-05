# chef-interpreter

[Chef](https://www.dangermouse.net/esoteric/chef.html) is an esoteric programming language developed by [David Morgan-Mar](https://www.dangermouse.net/). In Chef, programs are written to resemble real recipes as closely as possible.

`chef-interpreter` is a work in progress. Written in Go, it interprets recipes written in Chef and prints their results. Currently, it's only partially completed—stay tuned!

Two of my own Chef recipes are provided along with the interpreter for testing: Chili Weather and Caleb's Butterless Pi Crust. Chili Weather is an adaptation of my mom's own chili recipe—it takes a Celsius temperature as input and converts it to Fahrenheit. Caleb's Butterless Pi Crust is based on a pie crust recipe from a friend, and it calculates pi with six digits of precision. Each of these recipes is both executable and edible.

(For another Chef recipe, I highly recommend Mike Worth's [Hello World Cake with Chocolate Sauce](http://www.mike-worth.com/2013/03/31/baking-a-hello-world-cake/). This is one of my favorite recipes in Chef—not only does it print "Hello World," but like the other two, it also works as a real recipe.)
