# CVE-2021-21224 (issue 1195777)

> 1195777-chrome1day
https://www.youtube.com/embed/qZFzuHNqazo

http://127.0.0.1

## test

```
go run main.go
# visit http://localhost:3000/
```

## bugfix

+ src/v8/src/compiler/representation-change.cc

```c++
    if (output_type.Is(Type::Signed32()) ||
        output_type.Is(Type::Unsigned32())) {
      op = machine()->TruncateInt64ToInt32();
    } else if (output_type.Is(cache_->kSafeInteger) &&
               use_info.truncation().IsUsedAsWord32()) {
      op = machine()->TruncateInt64ToInt32();
    } else if (use_info.type_check() == TypeCheckKind::kSignedSmall ||
               use_info.type_check() == TypeCheckKind::kSigned32 ||
               use_info.type_check() == TypeCheckKind::kArrayIndex) {
      if (output_type.Is(cache_->kPositiveSafeInteger)) {
        op = simplified()->CheckedUint64ToInt32(use_info.feedback());
      } else if (output_type.Is(cache_->kSafeInteger)) {
        op = simplified()->CheckedInt64ToInt32(use_info.feedback());
      } else {
        return TypeError(node, output_rep, output_type,
                         MachineRepresentation::kWord32);
      }
    } else {
      return TypeError(node, output_rep, output_type,
                       MachineRepresentation::kWord32);
    }
```

```c++
    // BUGFIX: Issue 1195777 or CVE-2021-21224
    if (output_type.Is(Type::Signed32()) ||
        (output_type.Is(Type::Unsigned32()) &&
         use_info.type_check() == TypeCheckKind::kNone) ||
        (output_type.Is(cache_->kSafeInteger) &&
         use_info.truncation().IsUsedAsWord32())) {
      op = machine()->TruncateInt64ToInt32();
    } else if (use_info.type_check() == TypeCheckKind::kSignedSmall ||
               use_info.type_check() == TypeCheckKind::kSigned32 ||
               use_info.type_check() == TypeCheckKind::kArrayIndex) {
      if (output_type.Is(cache_->kPositiveSafeInteger)) {
        op = simplified()->CheckedUint64ToInt32(use_info.feedback());
      } else if (output_type.Is(cache_->kSafeInteger)) {
        op = simplified()->CheckedInt64ToInt32(use_info.feedback());
      } else {
        return TypeError(node, output_rep, output_type,
                         MachineRepresentation::kWord32);
      }
    } else {
      return TypeError(node, output_rep, output_type,
                       MachineRepresentation::kWord32);
    }
```
