// Golang Error Handling Is Better Than You Think!

package main

func main() {
	if err := userJustShitTheirPants(); err != nil {
		changeDaipers()
	}

	if err := bar(); err != nil {
		//handler
	}
}
func Foo() error {
	return nil
}
func bar() error {
	if err := Foo(); err != nil {
		return err
	}
	if err := Foo(); err != nil {
		return err
	}
	if err := Foo(); err != nil {
		return err
	}
	return nil
}
func changeDaipers() {

}

func userJustFarted() (int, error) {
	return 10, nil
}
func userJustShitTheirPants() error {
	return nil
}
