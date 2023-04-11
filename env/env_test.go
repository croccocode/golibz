package env

import (
	"errors"
	"testing"
)

func TestGetEnv(t *testing.T) {
	pwd, err := GetEnv("PWD")
	if err != nil {
		t.Error(err)
	}

	pwddie := GetenvOrDie("PWD")
	if pwddie != pwd {
		t.Errorf("excpected %s gto %s", pwd, pwddie)
	}

	pwdDefault := GetenvOrDefault("PWD", "Tintoriaaaa")
	if pwdDefault != pwd {
		t.Errorf("excpected %s gto %s", pwd, pwdDefault)
	}

	bestPodcastEver := GetenvOrDefault("NANANANANANANANABATMAAAAN", "Tintoriaaaa")
	if bestPodcastEver != "Tintoriaaaa" {
		t.Errorf("excpected 'Tintoriaaaa' gto %s", bestPodcastEver)
	}
}

func TestPanicIfErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	PanicIfErr(errors.New("top"))
}

func TestNoErrNoPanic(t *testing.T) {
	PanicIfErr(nil)
}
