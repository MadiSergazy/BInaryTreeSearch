package Web

import (
	"golangify.com/snippetbox/pkg/models"
	"html/template"
	"path/filepath"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func NewTemplateCache(dir string) (map[string]*template.Template, error) {

	cashe := map[string]*template.Template{}
	// Используем функцию filepath.Glob, чтобы получить срез всех файловых путей с
	// расширением '.page.tmpl'. По сути, мы получим список всех файлов шаблонов для страниц
	// нашего веб-приложения.
	pages, err := filepath.Glob(filepath.Join(dir, "*page.tmpl")) //home + show
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		//	Извлечение конечное названия файла (например, 'home.page.tmpl') из полного пути к файлу
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Используем метод ParseGlob для добавления всех каркасных шаблонов.
		// В нашем случае это только файл base.layout.tmpl (основная структура шаблона).
		ts, err = template.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// Используем метод ParseGlob для добавления всех вспомогательных шаблонов.
		// В нашем случае это footer.partial.tmpl "подвал" нашего шаблона.
		ts, err = template.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cashe[name] = ts
	}
	return cashe, nil
}
