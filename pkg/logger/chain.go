package logger

// Builder functions
func (l *Logger) Prefix(p string) *Logger {
	l.prefix = p
	return l
}

func (l *Logger) MaxLines(m int) *Logger {
	l.maxLines = m
	return l
}

func (l *Logger) STDOUT(b bool) *Logger {
	l.toStdout = b
	return l
}

func (l *Logger) JSONMode(b bool) *Logger {
	l.jsonMode = b
	return l
}

func (l *Logger) Debugging(b bool) *Logger {
	l.debugging = b
	return l
}

func (l *Logger) Subscribable(b bool) *Logger {
	l.subscribable = b
	if l.subscribable {
		l.subscribers = &Subscribers{}
	}

	return l
}

// Logging functions
func (l *Logger) Info(msg string) *Logger {
	l.log("info", msg)

	return l
}

func (l *Logger) Debug(msg string) *Logger {
	if l.debugging {
		l.log("debug", msg)
	}

	return l
}

func (l *Logger) Success(msg string) *Logger {
	l.log("success", msg)

	return l
}

func (l *Logger) Warning(msg string) *Logger {
	l.log("warning", msg)

	return l
}

func (l *Logger) Error(msg string) *Logger {
	l.log("error", msg)

	return l
}

func (l *Logger) Fatal(e error) *Logger {
	l.log("error", e.Error())

	return l
}
