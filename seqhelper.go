package go2linq

import (
	"context"
	"fmt"
	"iter"
	"strings"

	"github.com/solsw/errorhelper"
	"golang.org/x/sync/errgroup"
)

// SliceToSeq returns an iterator over elements of the [slice].
// If 's' is nil, empty iterator is returned.
//
// [slice]: https://go.dev/ref/spec#Slice_types
func SliceToSeq[S ~[]E, E any](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, e := range s {
			if !yield(e) {
				return
			}
		}
	}
}

// VarToSeq returns an iterator over elements of the [variadic] parameter.
//
// [variadic]: https://go.dev/ref/spec#Function_types
func VarToSeq[E any](s ...E) iter.Seq[E] {
	return SliceToSeq(s)
}

// SliceToSeq2 returns an iterator over pairs of index and value of elements of the [slice].
// If 's' is nil, empty iterator is returned.
//
// [slice]: https://go.dev/ref/spec#Slice_types
func SliceToSeq2[S ~[]E, E any](s S) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i, e := range s {
			if !yield(i, e) {
				return
			}
		}
	}
}

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
		return errorhelper.CallerError(ErrNilSource)
	}
	if action == nil {
		return errorhelper.CallerError(ErrNilAction)
	}
	for t := range seq {
		select {
		case <-ctx.Done():
			return errorhelper.CallerError(ctx.Err())
		default:
			if err := action(t); err != nil {
				return errorhelper.CallerError(err)
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
		return errorhelper.CallerError(ErrNilSource)
	}
	if action == nil {
		return errorhelper.CallerError(ErrNilAction)
	}
	g := new(errgroup.Group)
	for t := range seq {
		g.Go(func() error {
			select {
			case <-ctx.Done():
				return errorhelper.CallerError(ctx.Err())
			default:
				if err := action(t); err != nil {
					return errorhelper.CallerError(err)
				}
			}
			return nil
		})
	}
	return errorhelper.CallerError(g.Wait())
}

// SeqString converts a sequence to a sequence of strings.
func SeqString[T any](seq iter.Seq[T]) (iter.Seq[string], error) {
	return Select(seq, func(t T) string { return fmt.Sprint(t) })
}

// Strings returns a sequence contents as a slice of strings.
func Strings[T any](seq iter.Seq[T]) ([]string, error) {
	seqString, err := SeqString(seq)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	ss, err := ToSlice(seqString)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return ss, nil
}
