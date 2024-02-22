package go2linq

import (
	"context"
	"fmt"
	"iter"
	"strings"

	"golang.org/x/sync/errgroup"
)

// StringFmt returns string representation of a sequence:
//   - if 'seq' is nil, empty string is returned;
//   - if 'T' implements [fmt.Stringer], it is used to convert each element to string;
//   - 'sep' separates elements;
//   - 'lrim' and 'rrim' surround each element;
//   - 'ledge' and 'redge' surround the whole string.
func StringFmt[T any](seq iter.Seq[T], sep, lrim, rrim, ledge, redge string) string {
	if seq == nil {
		return ""
	}
	var b strings.Builder
	for t := range seq {
		if b.Len() > 0 {
			b.WriteString(sep)
		}
		b.WriteString(lrim + fmt.Sprint(t) + rrim)
	}
	return ledge + b.String() + redge
}

// StringDef returns string representation of a sequence using default formatting.
// If 'seq' is nil, empty string is returned.
func StringDef[T any](seq iter.Seq[T]) string {
	if seq == nil {
		return ""
	}
	return StringFmt(seq, " ", "", "", "[", "]")
}

// StringFmt2 returns string representation of a sequence:
//   - if 'seq2' is nil, empty string is returned;
//   - if 'T' implements [fmt.Stringer], it is used to convert each element to string;
//   - 'psep' separates pair of values;
//   - 'esep' separates elements;
//   - 'lrim' and 'rrim' surround each element;
//   - 'ledge' and 'redge' surround the whole string.
func StringFmt2[K, V any](seq2 iter.Seq2[K, V], psep, esep, lrim, rrim, ledge, redge string) string {
	if seq2 == nil {
		return ""
	}
	var b strings.Builder
	for k, v := range seq2 {
		if b.Len() > 0 {
			b.WriteString(esep)
		}
		b.WriteString(lrim + fmt.Sprint(k) + psep + fmt.Sprint(v) + rrim)
	}
	return ledge + b.String() + redge
}

// StringDef2 returns string representation of a sequence using default formatting.
// If 'seq2' is nil, empty string is returned.
func StringDef2[K, V any](seq2 iter.Seq2[K, V]) string {
	if seq2 == nil {
		return ""
	}
	return StringFmt2(seq2, ":", " ", "", "", "[", "]")
}

// ForEach sequentially performs a specified 'action' on each element of the sequence.
// If 'ctx' is canceled or 'action' returns non-nil error,
// operation is stopped and corresponding error is returned.
func ForEach[T any](ctx context.Context, seq iter.Seq[T], action func(T) error) error {
	if seq == nil {
		return ErrNilSource
	}
	if action == nil {
		return ErrNilAction
	}
	for t := range seq {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := action(t); err != nil {
				return err
			}
		}
	}
	return nil
}

// ForEachConcurrent concurrently performs a specified 'action' on each element of the sequence.
// If 'ctx' is canceled or 'action' returns non-nil error,
// operation is stopped and corresponding error is returned.
func ForEachConcurrent[T any](ctx context.Context, seq iter.Seq[T], action func(T) error) error {
	if seq == nil {
		return ErrNilSource
	}
	if action == nil {
		return ErrNilAction
	}
	g := new(errgroup.Group)
	for t := range seq {
		g.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				if err := action(t); err != nil {
					return err
				}
			}
			return nil
		})
	}
	return g.Wait()
}

// SeqString converts a sequence to a sequence of strings.
func SeqString[T any](seq iter.Seq[T]) (iter.Seq[string], error) {
	return Select[T, string](seq, func(t T) string { return fmt.Sprint(t) })
}

// Strings returns a sequence contents as a slice of strings.
func Strings[T any](seq iter.Seq[T]) ([]string, error) {
	seqString, err := SeqString[T](seq)
	if err != nil {
		return nil, err
	}
	ss, err := ToSlice(seqString)
	if err != nil {
		return nil, err
	}
	return ss, nil
}
