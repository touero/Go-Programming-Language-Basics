package ledger

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

var errInvalid = errors.New("invalid input")

// localeConfig 描述一个 locale 的表头、日期格式、金额格式
type localeConfig struct {
	header       string
	formatDate   func(y, m, d string) string
	formatAmount func(currency string, cents int) (string, error)
}

var locales = map[string]localeConfig{
	"en-US": {
		header: "Date" + pad("Date", 10) + " | " +
			"Description" + pad("Description", 25) + " | " +
			"Change" + pad("Change", 13),
		formatDate: func(y, m, d string) string {
			return m + "/" + d + "/" + y
		},
		formatAmount: formatAmountEnUS,
	},
	"nl-NL": {
		header: "Datum" + pad("Datum", 10) + " | " +
			"Omschrijving" + pad("Omschrijving", 25) + " | " +
			"Verandering" + pad("Verandering", 13),
		formatDate: func(y, m, d string) string {
			return d + "-" + m + "-" + y
		},
		formatAmount: formatAmountNlNL,
	},
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	cfg, ok := locales[locale]
	if !ok {
		return "", errInvalid
	}

	// 先校验 currency，空 entries 时也要报错
	if _, err := symbol(currency); err != nil {
		return "", err
	}

	// 防御性复制 + 排序（不修改调用方切片）
	sorted := make([]Entry, len(entries))
	copy(sorted, entries)
	sort.SliceStable(sorted, func(i, j int) bool {
		a, b := sorted[i], sorted[j]
		if a.Date != b.Date {
			return a.Date < b.Date
		}
		if a.Description != b.Description {
			return a.Description < b.Description
		}
		return a.Change < b.Change
	})

	var sb strings.Builder
	sb.WriteString(cfg.header)
	sb.WriteString("\n")

	for _, e := range sorted {
		line, err := formatEntry(cfg, currency, e)
		if err != nil {
			return "", err
		}
		sb.WriteString(line)
		sb.WriteString("\n")
	}
	return sb.String(), nil
}

func formatEntry(cfg localeConfig, currency string, e Entry) (string, error) {
	// 校验并拆分日期
	if len(e.Date) != 10 || e.Date[4] != '-' || e.Date[7] != '-' {
		return "", errInvalid
	}
	y, m, d := e.Date[0:4], e.Date[5:7], e.Date[8:10]
	dateStr := cfg.formatDate(y, m, d)

	// 描述截断/补齐（保持原代码按字节的行为）
	desc := e.Description
	if len(desc) > 25 {
		desc = desc[:22] + "..."
	} else {
		desc = desc + strings.Repeat(" ", 25-len(desc))
	}

	// 金额
	amount, err := cfg.formatAmount(currency, e.Change)
	if err != nil {
		return "", err
	}

	// 金额列宽 13，按 rune 计长（金额含 €/$ 等多字节）
	amountWidth := utf8RuneCount(amount)
	pad := 13 - amountWidth
	if pad < 0 {
		pad = 0
	}

	return dateStr + strings.Repeat(" ", 10-len(dateStr)) +
		" | " + desc + " | " +
		strings.Repeat(" ", pad) + amount, nil
}

func pad(s string, width int) string {
	return strings.Repeat(" ", width-len(s))
}

func utf8RuneCount(s string) int {
	n := 0
	for range s {
		n++
	}
	return n
}

func formatAmountEnUS(currency string, cents int) (string, error) {
	sym, err := symbol(currency)
	if err != nil {
		return "", err
	}
	neg := cents < 0
	if neg {
		cents = -cents
	}
	intPart, frac := splitCents(cents)
	s := sym + groupThousands(intPart, ",") + "." + frac
	if neg {
		s = "(" + s + ")" // 原行为：负数右括号后不加空格
	} else {
		s += " "
	}
	return s, nil
}

func formatAmountNlNL(currency string, cents int) (string, error) {
	sym, err := symbol(currency)
	if err != nil {
		return "", err
	}
	neg := cents < 0
	if neg {
		cents = -cents
	}
	intPart, frac := splitCents(cents)
	s := sym + " "
	if neg {
		s += "-"
	}
	s += groupThousands(intPart, ".") + "," + frac + " "
	return s, nil
}

func symbol(currency string) (string, error) {
	switch currency {
	case "EUR":
		return "€", nil
	case "USD":
		return "$", nil
	default:
		return "", errInvalid
	}
}

// splitCents 把分拆成整数部分字符串和两位小数
func splitCents(cents int) (intPart, frac string) {
	s := strconv.Itoa(cents)
	for len(s) < 3 {
		s = "0" + s
	}
	return s[:len(s)-2], s[len(s)-2:]
}

// groupThousands 从右往左每 3 位插分隔符
func groupThousands(intPart, sep string) string {
	if len(intPart) <= 3 {
		return intPart
	}
	var parts []string
	for len(intPart) > 3 {
		parts = append(parts, intPart[len(intPart)-3:])
		intPart = intPart[:len(intPart)-3]
	}
	parts = append(parts, intPart)
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	return strings.Join(parts, sep)
}
