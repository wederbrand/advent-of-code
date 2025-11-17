#!/usr/bin/env perl
use strict;
use warnings;

# Open input file
open my $fh, '<', 'input.txt' or die "Cannot open input.txt: $!";

# Read one line (strip newline)
my $line = <$fh>;
chomp $line if defined $line;
close $fh;

# Exit if no line
exit 0 unless defined $line && length $line;

# Split by comma to get pairs
my @pairs = split /,/, $line;


my $result = 0;
for my $pair (@pairs) {
    # Each pair is start-end
    my ($start, $end) = split /-/, $pair;

    # Skip malformed pairs
    next unless defined $start && defined $end;
    $start += 0;  # force numeric
    $end   += 0;

    # Iterate from start to end (inclusive)
    if ($start <= $end) {
        for my $n ($start .. $end) {
            if ($n =~ /^(.+)\1+$/) {
                $result+=$n;
            }
        }
    }
}

# Print the result
print "$result\n";