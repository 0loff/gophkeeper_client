// Пакет для логирования в приложении.
package logger

import (
	"go.uber.org/zap"
)

// Конструктор инициализации логгера Zap
var Log *zap.Logger = zap.NewNop()

// Упрощенный метод вызова ллоггера Zap с помощью синтаксического сахара
var Sugar *zap.SugaredLogger

// Инициализация механизма логирования в приложении
func Initialize(level string) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = lvl

	zl, err := cfg.Build()
	if err != nil {
		return err
	}

	Log = zl
	Sugar = Log.Sugar()

	return nil
}
